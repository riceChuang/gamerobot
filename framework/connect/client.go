package connect

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sync"
)

// Client wrapper ws and decoder
type Client struct {
	conn            *Conn
	parser          model.Parser
	evtChan         *eventChan
	innerHandler    sync.Map
	onCloseDelegate func() // if socket dead by read / dispatch / write error, please take error from here
	unHandleBytes   []byte
	lock            sync.Mutex // readwrite lock
	logger          *log.Entry
}

// NewClient new a net client
func NewClient(conn *Conn, parser model.Parser) *Client {
	nc := &Client{
		conn:   conn,
		parser: parser,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "client",
		}),
	}
	// nc.innerHandler = make(map[string]*msg.Handler)
	nc.innerHandler = sync.Map{}
	nc.evtChan = &eventChan{
		ReadSig:  make(chan *model.Message, 2048),
		WriteSig: make(chan *model.Message, 2048),
		ErrSig:   make(chan error),
		StopSig:  make(chan bool, 1),
	}



	return nc
}

// Connect ...
func (nc *Client) Connect() error {
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
func (nc *Client) Write(msg *model.Message) {
	nc.evtChan.WriteSig <- msg
}

// Close ...
func (nc *Client) Close() error {
	nc.evtChan.StopSig <- true
	if err := nc.conn.Close(); err != nil {
		return err
	}
	return nil
}

// Clean ...
func (nc *Client) Clean() {
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
func (nc *Client) Register(h *model.Handler) error {
	if _, exist := nc.innerHandler.Load(h.String()); exist {
		return fmt.Errorf("handler already exist: %s", h.String())
	}
	nc.innerHandler.Store(h.String(), h)
	return nil
}

// UnRegister Handler ...
func (nc *Client) UnRegister(h *model.Handler) {
	nc.innerHandler.Delete(h.String())
}

func (nc *Client) read() {
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

			message, err := nc.parser.UnMarshal(tempBytes[:length+4])
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

	log.Info("[ERROR]  SOCKET READ BREAK ")
}

func (nc *Client) handMsg() {
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

func (nc *Client) write() {
ForWrite:
	for {
		select {
		case message := <-nc.evtChan.WriteSig:
			bytes, err := nc.parser.Marshal(message)
			if err != nil {
				log.Errorf("Marshal Error, %v", err.Error())
				nc.evtChan.ErrSig <- err
				break ForWrite
			}
			//nc.iLog.Println("Client Write", message.GetName())
			err = nc.conn.Write(bytes)
			if err != nil {
				log.Errorf("Write Error, %v", err.Error())
				nc.evtChan.ErrSig <- err
				break ForWrite
			}
		case <-nc.evtChan.StopSig:
			break ForWrite
		}
	}
}

func (nc *Client) dispatch(m *model.Message) {
	if data, ok := nc.innerHandler.Load(m.String()); ok {
		h := data.(*model.Handler)
		log.Infof("dispatch message, %v", m.GetName())
		if h.MsgType != nil {
			if err := proto.Unmarshal(m.Data.([]byte), h.MsgType); err != nil {
				log.Errorf("client proto unmarshal error: ", err, fmt.Sprintf("%+v", m.Data))
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
func (nc *Client) onClose(err error) {
	// log.Printf("[DEBUG][onCLOSE] ---- onClose %v", err)
	if nc.conn != nil {
		if err != nil {
			log.Panicf("Client Close By err", err.Error())
		}
		if nc.onCloseDelegate != nil {
			nc.onCloseDelegate()
		}
		nc.Close()
	}
}
