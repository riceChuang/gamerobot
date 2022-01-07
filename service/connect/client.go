package connect

import (
	"context"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/connection/connect"
	"github.com/riceChuang/gamerobot/framework/connection/connecttype"
	"github.com/riceChuang/gamerobot/game/gametype"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ClientConn struct {
	ctx      context.Context
	ID       string
	wsConn   *connect.HttpConnect
	gameConn *connecttype.GameConnect
	GameID   common.GameServerID
	GameType gametype.GameType
	Parser   util.Parser
	logger   *log.Entry
	lock     sync.Mutex // readwrite lock
}

func NewClientWsGameConn(ctx context.Context) (*ClientConn, error) {
	client := &ClientConn{
		ID:  uuid.NewV4().String(),
		ctx: ctx,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "client_conn",
		}),
		Parser: &util.ByteParser{},
	}
	return client, nil
}

func (client *ClientConn) SetWsConn(connect *connect.HttpConnect) {
	connect.ClientID = client.ID
	client.wsConn = connect
	return
}

func (client *ClientConn) SetGameConn(connect *connecttype.GameConnect) {
	connect.ClientID = client.ID
	client.gameConn = connect
	return
}

func (client *ClientConn) SetGameID(id common.GameServerID) {
	client.GameID = id
	return
}

func (client *ClientConn) SetGameType(gameType gametype.GameType) {
	client.GameType = gameType
	return
}

func (client *ClientConn) ProtoConnect() error {
	if client.GameType != nil {
		client.GameType.Initialize(client.gameConn)
	}
	client.gameConn.ConnectGameWs()
	return nil
}

func (client *ClientConn) IsGameConnAlive() bool {
	if client.GameID != 0 && client.GameType != nil && client.gameConn != nil {
		return true
	}
	return false
}

func (client *ClientConn) IsWsConnAlive() bool {
	return client.wsConn != nil
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
