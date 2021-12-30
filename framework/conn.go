package framework

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"gitlab.baifu-tech.net/dsg-game/game-robot/util"
	"gitlab.baifu-tech.net/dsg-game/game-robot/util/logs"
	"io"
	"sync"
)

// Conn handle read and write, also connection
type Conn struct {
	URL       string          // 連接網址
	Conn      *websocket.Conn // Conn實體
	IsConnect bool            // 是否連接
	lock      sync.Mutex      // readwrite lock
	logger    *log.Entry      // inner Logger
	wsType    int
}

// NewConn new a ws Conn
func NewConn(URL string, conn *websocket.Conn) *Conn {
	c := &Conn{
		URL: URL,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "client_conn",
		}),
		wsType: websocket.TextMessage,
	}
	if conn != nil {
		c.Conn = conn
		c.IsConnect = true
	}
	return c
}

// Connect connect to url
func (ws *Conn) Connect() error {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	if !ws.IsConnect {
		conn, _, err := websocket.DefaultDialer.Dial(ws.URL, nil)
		if err != nil {
			return fmt.Errorf("socket dial error: %s", err.Error())
		}
		ws.Conn = conn
		ws.IsConnect = true
		ws.logger.Println("Socket Connection success", ws.URL)
	} else {
		return errors.New("socket already connect")
	}
	return nil
}

// Read from Socket
func (ws *Conn) Read() ([]byte, error) {
	if ws.IsConnect {
		_, bytes, err := ws.Conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				return nil, err
			}
			return nil, fmt.Errorf("Read Error: %s", err.Error())
		}
		return bytes, nil
	}
	return nil, errors.New("socket Read Error: You need to Connect first")
}

// Write to Socket
func (ws *Conn) Write(bytes []byte) error {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	if ws.IsConnect {
		length := util.IntToBytes(int32(len(bytes)), true)
		// fmt.Println("send ..... ", length, append(length, bytes...))
		ws.Conn.WriteMessage(ws.wsType, append(length, bytes...))
	} else {
		return errors.New("socket Write Error: You need to Connect first")
	}
	return nil
}

// Close socket
func (ws *Conn) Close() error {
	ws.lock.Lock()
	defer ws.lock.Unlock()
	if ws.IsConnect {
		err := ws.Conn.Close()
		if err != nil {
			return fmt.Errorf("socket Close Error: %s", err.Error())
		}
		ws.IsConnect = false
		ws.Conn = nil
	} else {
		return errors.New("socket already closed")
	}
	return nil
}

func (ws *Conn) SetWsType(wsType int) {
	ws.wsType = wsType
}
