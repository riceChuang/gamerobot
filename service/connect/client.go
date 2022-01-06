package connect

import (
	"context"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/util/logs"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ClientConn struct {
	ctx       context.Context
	ID        string
	wsConn    *framework.HttpConnect
	gameConn  *GameConnect
	Parser    framework.Parser
	GameID    common.GameServerID
	logger    *log.Entry
	lock      sync.Mutex // readwrite lock
}

func NewClientWsGameConn(ctx context.Context) (*ClientConn, error) {
	client := &ClientConn{
		ID:  uuid.NewV4().String(),
		ctx: ctx,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "client_ws",
		}),
	}
	return client, nil
}

func (client *ClientConn) SetWsConn(connect *framework.HttpConnect) {
	connect.ClientID = client.ID
	client.wsConn = connect
}

func (client *ClientConn) SetGameConn(connect *GameConnect) {
	connect.ClientID = client.ID
	client.gameConn = connect
}

func (client *ClientConn) SetGameID(id common.GameServerID) {
	client.GameID = id
}

func (client *ClientConn) ProtoConnect() error {
	client.gameConn.ConnectGameWs()
	return nil
}

//收clinet ws 資訊
//func (conn *ClientConn) ReadClientWS() ([]byte, error) {
//	if conn.wsConn == nil || conn.wsConn.Conn == nil {
//		return nil, errors.New("找不到ws conn")
//	}
//	return conn.wsConn.Read()
//}
//
////寫入client ws 資訊
//func (conn *ClientConn) WriteClientWS(bytes []byte) error {
//	if conn.wsConn == nil || conn.wsConn.Conn == nil {
//		return errors.New("找不到ws conn")
//	}
//	return conn.wsConn.Write(bytes)
//}
//
////釋放client ws
//func (conn *ClientConn) ReleaseClientWs() error {
//	return conn.wsConn.Close()
//}
//
////收clinet ws 資訊
//func (conn *ClientConn) ReadGameWS() ([]byte, error) {
//	if conn.gameConn == nil || conn.gameConn.Conn == nil {
//		return nil, errors.New("找不到ws conn")
//	}
//	return conn.gameConn.Read()
//}

////寫入client ws 資訊
//func (conn *ClientConn) WriteGameWS(bytes []byte) error {
//	if conn.gameConn == nil || conn.gameConn.Conn == nil {
//		return errors.New("找不到ws conn")
//	}
//	return conn.gameConn.Write(bytes)
//}
//
////釋放client ws
//func (conn *ClientConn) ReleaseGameWs() error {
//	return conn.gameConn.Close()
//}
