package connect

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	"sync"

	"github.com/golang/protobuf/proto"
)

// Client wrapper ws and decoder
type ProtoConnect struct {
	conn            *framework.Conn
	Parser          framework.Parser
	evtChan         *model.EventChan
	innerHandler    sync.Map
	onCloseDelegate func() // if socket dead by read / dispatch / write error, please take error from here
	unHandleBytes   []byte
	lock            sync.Mutex    // readwrite lock
	logger          *log.Entry // inner Logger
}

func NewProtoConnect(conn *framework.Conn) *ProtoConnect {
	nc := &ProtoConnect{
		conn:   conn,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "proto_client",
		}),
	}
	// nc.innerHandler = make(map[string]*msg.Handler)
	nc.innerHandler = sync.Map{}
	nc.evtChan = &model.EventChan{
		ReadSig:  make(chan *model.Message, 2048),
		WriteSig: make(chan *model.Message, 2048),
		ErrSig:   make(chan error),
		StopSig:  make(chan bool, 1),
	}
	return nc
}

// Connect ...
func (nc *ProtoConnect) Connect() error {
	err := nc.conn.Connect()
	if err != nil {
		return err
	}
	go nc.read()
	go nc.handMsg()
	go nc.write()
	return nil
}

// Write data with encode
func (nc *ProtoConnect) Write(msg *model.Message) {
	nc.evtChan.WriteSig <- msg
}

// Close ...
func (nc *ProtoConnect) Close() error {
	nc.evtChan.StopSig <- true
	if err := nc.conn.Close(); err != nil {
		return err
	}
	return nil
}

// Clean ...
func (nc *ProtoConnect) Clean() {
	// nc.iLog.Println("[DEBUG][Client][Clean] ----")
	nc.lock.Lock()
	defer nc.lock.Unlock()
	nc.onClose(nil)
	nc.innerHandler.Range(func(key, _ interface{}) bool {
		nc.innerHandler.Delete(key)
		return true
	})
	nc.conn = nil
}

// Register Handler ...
func (nc *ProtoConnect) Register(h *model.Handler) error {
	if _, exist := nc.innerHandler.Load(h.String()); exist {
		return fmt.Errorf("handler already exist: %s", h.String())
	}
	nc.logger.Info("Register handler", h.String())
	nc.innerHandler.Store(h.String(), h)
	return nil
}

// UnRegister Handler ...
func (nc *ProtoConnect) UnRegister(h *model.Handler) {
	nc.innerHandler.Delete(h.String())
}

func (nc *ProtoConnect) read() {
ForRead:
	for {
		data, err := nc.conn.Read()
		if err != nil {
			nc.evtChan.ErrSig <- err
			break ForRead
		}

		var tempBytes []byte
		if nc.unHandleBytes != nil {
			tempBytes = append(nc.unHandleBytes, data...)
			nc.unHandleBytes = nil
		} else {
			tempBytes = data
		}

		for len(tempBytes) >= 4 {
			// Force drop wired package
			length := util.BytesToInt(tempBytes[:4], true)
			// if int(length+4) > 10000 {
			// 	nc.iLog.Println("[DEUBG]強迫drop 異常包", int(length+4), len(tempBytes))
			// 	nc.unHandleBytes = nil
			// 	tempBytes = nil
			// 	continue
			// }

			if int(length+4) > len(tempBytes) {
				nc.unHandleBytes = tempBytes
				tempBytes = nil
				continue
			}

			message, err := nc.Parser.UnMarshal(tempBytes[:length+4])
			if err != nil {
				nc.evtChan.ErrSig <- err
				break ForRead
			}
			nc.evtChan.ReadSig <- message

			if int(length+4) == len(tempBytes) {
				tempBytes = nil
			} else {
				tempBytes = tempBytes[length+4:]
			}
		}
		if len(tempBytes) != 0 && len(tempBytes) < 4 {
			nc.unHandleBytes = tempBytes
		}
	}

	nc.logger.Println("[ERROR]  SOCKET READ BREAK ")
}

func (nc *ProtoConnect) handMsg() {
ForHandle:
	for {
		select {
		case err := <-nc.evtChan.ErrSig:
			nc.onClose(err)
		case message := <-nc.evtChan.ReadSig:
			nc.dispatch(message)
		case <-nc.evtChan.StopSig:
			break ForHandle
		}
	}
}

func (nc *ProtoConnect) write() {
ForWrite:
	for {
		select {
		case message := <-nc.evtChan.WriteSig:
			bytes, err := nc.Parser.Marshal(message)
			if err != nil {
				nc.logger.Println("Marshal Error", err.Error())
				nc.evtChan.ErrSig <- err
				break ForWrite
			}
			//nc.iLog.Println("Client Write", message.GetName())
			err = nc.conn.Write(bytes)
			if err != nil {
				nc.logger.Println("Write Error", err.Error())
				nc.evtChan.ErrSig <- err
				break ForWrite
			}
		case <-nc.evtChan.StopSig:
			break ForWrite
		}
	}
}

func (nc *ProtoConnect) dispatch(m *model.Message) {
	if data, ok := nc.innerHandler.Load(m.String()); ok {
		h := data.(*model.Handler)
		nc.logger.Println("dispatch message", m.GetName())
		if h.MsgType != nil {
			if err := proto.Unmarshal(m.Data.([]byte), h.MsgType); err != nil {
				nc.logger.Println("client proto unmarshal error: ", err.Error(), fmt.Sprintf("%+v", m.Data))
				nc.evtChan.ErrSig <- err
				return
			}
		}
		go h.OnMessage(h.MsgType)
	} else {
		// nc.iLog.Println("no hander exist drop: ", m.GetName())
	}
}

// onClose
func (nc *ProtoConnect) onClose(err error) {
	// log.Printf("[DEBUG][onCLOSE] ---- onClose %v", err)
	if nc.conn != nil {
		if err != nil {
			nc.logger.Panic("Client Close By err", err.Error())
		}
		if nc.onCloseDelegate != nil {
			nc.onCloseDelegate()
		}
		nc.Close()
	}
}
