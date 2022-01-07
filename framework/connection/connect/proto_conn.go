package connect

import (
	"fmt"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
)

// Client wrapper ws and decoder
type ProtoConnect struct {
	conn            *Conn
	Parser          util.Parser
	evtChan         *model.EventChan
	innerHandler    sync.Map
	onCloseDelegate func() // if socket dead by read / dispatch / write error, please take error from here
	unHandleBytes   []byte
	lock            sync.Mutex // readwrite lock
	logger          *log.Entry // inner Logger
	timerChan       chan bool
}

func NewProtoConnect(conn *Conn) *ProtoConnect {
	pc := &ProtoConnect{
		conn: conn,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "proto_connect",
		}),
		Parser: &util.ByteParser{},
	}
	// pc.innerHandler = make(map[string]*msg.Handler)
	pc.innerHandler = sync.Map{}
	pc.evtChan = &model.EventChan{
		ReadSig:  make(chan *model.Message, 2048),
		WriteSig: make(chan *model.Message, 2048),
		ErrSig:   make(chan error),
		StopSig:  make(chan bool, 1),
	}
	return pc
}

// Connect ...
func (pc *ProtoConnect) Connect() error {
	err := pc.conn.Connect()
	if err != nil {
		return err
	}
	go pc.read()
	go pc.handMsg()
	go pc.write()
	pc.initHeartBeat()
	return nil
}

// Write data with encode
func (pc *ProtoConnect) Write(msg *model.Message) {
	pc.evtChan.WriteSig <- msg
}

// Close ...
func (pc *ProtoConnect) Close() error {
	pc.evtChan.StopSig <- true
	if err := pc.conn.Close(); err != nil {
		return err
	}
	return nil
}

// Clean ...
func (pc *ProtoConnect) Clean() {
	// pc.iLog.Println("[DEBUG][Client][Clean] ----")
	pc.lock.Lock()
	defer pc.lock.Unlock()
	pc.onClose(nil)
	pc.innerHandler.Range(func(key, _ interface{}) bool {
		pc.innerHandler.Delete(key)
		return true
	})
	pc.conn = nil
}

// Register Handler ...
func (pc *ProtoConnect) Register(h *model.Handler) error {
	if _, exist := pc.innerHandler.Load(h.String()); exist {
		return fmt.Errorf("handler already exist: %s", h.String())
	}
	pc.logger.Info("Register handler", h.String())
	pc.innerHandler.Store(h.String(), h)
	return nil
}

// UnRegister Handler ...
func (pc *ProtoConnect) UnRegister(h *model.Handler) {
	pc.innerHandler.Delete(h.String())
}

func (pc *ProtoConnect) GetConnectURL() string {
	if pc.conn != nil {
		return pc.conn.URL
	}
	return ""
}

func (pc *ProtoConnect) initHeartBeat() {
	pc.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_Common),
		SClassID:  int32(netproto.PlatformCommonClassID_HeartBeatConfigID),
		MsgType:   &netproto.HeartBeatConfig{},
		OnMessage: pc.onHeartBeatCfg,
	})

	pc.requestHeartBeatCfg()
}

func (pc *ProtoConnect) requestHeartBeatCfg() {
	message, err := model.NewProto(
		int32(netproto.MessageBClassID_Common),
		int32(netproto.PlatformCommonClassID_RequestHeartBeatConfigID),
		nil,
	)

	if err != nil {
		pc.logger.Println("Error: send heart beat failed", err.Error())
	}
	pc.logger.Infof("[請求requestHeartBeatCfg]")
	pc.Write(message)
}

func (pc *ProtoConnect) onHeartBeatCfg(b interface{}) {
	data, ok := b.(*netproto.HeartBeatConfig)
	if !ok {
		// should kill robot
		pc.logger.Panic("Error: Data Type Wrong!")
		return
	}
	pc.logger.Infof("[收到requestHeartBeatCfg]")
	//pc.logger.Println("onHeartBeatCfg:", data.String())

	if pc.timerChan != nil {
		pc.timerChan <- true
	}
	pc.timerChan = util.NewTicker(pc.sendHeartBeat, time.Duration(data.GetBeatInterval())*time.Second)
}

func (pc *ProtoConnect) sendHeartBeat() {
	message, err := model.NewProto(
		int32(netproto.MessageBClassID_Common),
		int32(netproto.PlatformCommonClassID_HeartBeatID),
		nil,
	)

	if err != nil {
		pc.logger.Println("Error: send heart beat failed", err.Error())
	}
	//pc.logger.Debug("[發送心跳]")
	pc.Write(message)
}

func (pc *ProtoConnect) read() {
ForRead:
	for {
		data, err := pc.conn.Read()
		if err != nil {
			pc.evtChan.ErrSig <- err
			break ForRead
		}

		var tempBytes []byte
		if pc.unHandleBytes != nil {
			tempBytes = append(pc.unHandleBytes, data...)
			pc.unHandleBytes = nil
		} else {
			tempBytes = data
		}

		for len(tempBytes) >= 4 {
			// Force drop wired package
			length := util.BytesToInt(tempBytes[:4], true)
			// if int(length+4) > 10000 {
			// 	pc.iLog.Println("[DEUBG]強迫drop 異常包", int(length+4), len(tempBytes))
			// 	pc.unHandleBytes = nil
			// 	tempBytes = nil
			// 	continue
			// }

			if int(length+4) > len(tempBytes) {
				pc.unHandleBytes = tempBytes
				tempBytes = nil
				continue
			}

			message, err := pc.Parser.UnMarshal(tempBytes[:length+4])
			if err != nil {
				pc.evtChan.ErrSig <- err
				break ForRead
			}
			pc.evtChan.ReadSig <- message

			if int(length+4) == len(tempBytes) {
				tempBytes = nil
			} else {
				tempBytes = tempBytes[length+4:]
			}
		}
		if len(tempBytes) != 0 && len(tempBytes) < 4 {
			pc.unHandleBytes = tempBytes
		}
	}

	pc.logger.Println("[ERROR]  SOCKET READ BREAK ")
}

func (pc *ProtoConnect) handMsg() {
ForHandle:
	for {
		select {
		case err := <-pc.evtChan.ErrSig:
			pc.onClose(err)
		case message := <-pc.evtChan.ReadSig:
			pc.dispatch(message)
		case <-pc.evtChan.StopSig:
			break ForHandle
		}
	}
}

func (pc *ProtoConnect) write() {
ForWrite:
	for {
		select {
		case message := <-pc.evtChan.WriteSig:
			bytes, err := pc.Parser.Marshal(message)
			if err != nil {
				pc.logger.Println("Marshal Error", err.Error())
				pc.evtChan.ErrSig <- err
				break ForWrite
			}
			//pc.iLog.Println("Client Write", message.GetName())
			err = pc.conn.Write(bytes)
			if err != nil {
				pc.logger.Println("Write Error", err.Error())
				pc.evtChan.ErrSig <- err
				break ForWrite
			}
		case <-pc.evtChan.StopSig:
			break ForWrite
		}
	}
}

func (pc *ProtoConnect) dispatch(m *model.Message) {
	pc.logger.Infof("dispatch Message :%v", m.GetName())
	if data, ok := pc.innerHandler.Load(fmt.Sprintf("%d:0", m.BClassID)); ok {
		if innerData, k := pc.innerHandler.Load(m.String()); k {
			h := innerData.(*model.Handler)
			pc.sendMessage(m, h)
			return
		}
		//若沒有指定則送Bclass層級
		h := data.(*model.Handler)
		pc.sendMessage(m, h)
		return
	}

	if innerData, k := pc.innerHandler.Load(m.String()); k {
		h := innerData.(*model.Handler)
		pc.sendMessage(m, h)
		return
	}
}

func (pc *ProtoConnect) sendMessage(m *model.Message, h *model.Handler) {
	pc.logger.Info("dispatch message", m.GetName())
	if h.MsgType != nil {
		if err := proto.Unmarshal(m.Data.([]byte), h.MsgType); err != nil {
			pc.logger.Println("client proto unmarshal error: ", err.Error(), fmt.Sprintf("%+v", m.Data))
			pc.evtChan.ErrSig <- err
			return
		}
		h.OnMessage(h.MsgType)
		return
	} else {
		h.OnMessage(m)
	}
}

// onClose
func (pc *ProtoConnect) onClose(err error) {
	// log.Printf("[DEBUG][onCLOSE] ---- onClose %v", err)
	if pc.conn != nil {
		if err != nil {
			pc.logger.Panic("Client Close By err", err.Error())
		}
		if pc.onCloseDelegate != nil {
			pc.onCloseDelegate()
		}
		pc.Close()
	}
}
