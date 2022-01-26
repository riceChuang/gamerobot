package connect

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	"sync"
	"time"
)

const (
	// Time allowed to read the next pong message from the peer.
	pingWait = 60 * time.Second
)

// Client wrapper ws and decoder
type HttpConnect struct {
	conn         *Conn
	innerHandler sync.Map
	lock         sync.Mutex // readwrite lock
	logger       *log.Entry // inner Logger
	ClientID     string
}

func NewHttpConnect(conn *Conn) *HttpConnect {
	hc := &HttpConnect{
		conn: conn,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "http_connect",
		}),
	}
	// nc.innerHandler = make(map[string]*msg.Handler)
	hc.innerHandler = sync.Map{}
	return hc
}

// Connect ...
func (hc *HttpConnect) Connect() error {
	go hc.read()
	return nil
}

func (hc *HttpConnect) Write(msg *model.Message) {
	//sendMsg, err := json.Marshal(msg)
	//if err != nil {
	//	hc.logger.Error("Marshal err:%v", err)
	//	return
	//}
	err := hc.conn.Write(msg.ToStringMessage())
	if err != nil {
		hc.logger.Error("資料錯誤:%v", err)
		return
	}
}

// Close ...
func (hc *HttpConnect) Close() error {
	if hc.conn == nil {
		return nil
	}
	if err := hc.conn.Close(); err != nil {
		return err
	}
	hc.conn = nil
	return nil
}

// Clean ...
func (hc *HttpConnect) Clean() {
	// nc.iLog.Println("[DEBUG][Client][Clean] ----")
	hc.lock.Lock()
	defer hc.lock.Unlock()
	hc.innerHandler.Range(func(key, _ interface{}) bool {
		hc.innerHandler.Delete(key)
		return true
	})
	hc.conn = nil
}

// Register Handler ...
func (nc *HttpConnect) Register(h *model.Handler) error {
	if _, exist := nc.innerHandler.Load(h.String()); exist {
		return fmt.Errorf("handler already exist: %s", h.String())
	}
	nc.logger.Info("Register handler", h.String())
	nc.innerHandler.Store(h.String(), h)
	return nil
}

// UnRegister Handler ...
func (nc *HttpConnect) UnRegister(h *model.Handler) {
	nc.innerHandler.Delete(h.String())
}

func (hc *HttpConnect) read() {
	hc.readDeadLine(pingWait)
	for {
		defer hc.conn.Close()
		//读取ws中的数据
		message, err := hc.conn.Read()
		if err != nil {
			break
		}

		msg := &model.Message{}
		err = json.Unmarshal(message, msg)
		if err != nil {
			fmt.Println(err)
		}
		msg.Data = []byte{}
		if string(message) == "ping" {
			p :=hc.conn.Conn.PongHandler()
			err = p("pong")
			if err != nil {
				break
			}
			continue
		}

		wsMessage := &model.WSMessage{
			From:     common.Client,
			To:       common.ClientServerTransfer,
			ClientID: hc.ClientID,
			Msg:      msg,
		}
		util.GetClientDispatcher().AddMessage(wsMessage)
	}
	return
}


func (hc *HttpConnect) readDeadLine(wait time.Duration) {
	timmer := time.NewTimer(wait)
	hc.conn.Conn.SetPongHandler(func(string) error {
		timmer.Reset(wait)
		if err := hc.conn.Conn.WriteMessage(websocket.PongMessage, nil); err != nil {
			return err
		}
		return nil
	})

	go func () {
		for {
			select{
			case <-timmer.C :

				wsMessage := &model.WSMessage{
					From:     common.Client,
					To:       common.ClientServerTransfer,
					Msg:      &model.Message{
						BClassID: common.WSClose,
						Data: "",
					},
				}
				util.GetClientDispatcher().AddMessage(wsMessage)
				hc.Close()
				return
			}
		}
	}()
}