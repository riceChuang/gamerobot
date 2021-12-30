package connect

import (
	"context"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"gitlab.baifu-tech.net/dsg-game/game-robot/common"
	"gitlab.baifu-tech.net/dsg-game/game-robot/framework"
	"gitlab.baifu-tech.net/dsg-game/game-robot/util/logs"
	"sync"
)

type ClientConn struct {
	ctx              context.Context
	ID               string
	wsConn           *HttpConnect
	gameConn         *ProtoConnect
	Parser           framework.Parser
	GameID           common.GameServerID
	logger           *log.Entry
	lock             sync.Mutex // readwrite lock
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

func (conn *ClientConn) SetWsConn(connect *HttpConnect) {
	connect.ClientID = conn.ID
	conn.wsConn = connect
}

func (conn *ClientConn) SetGameConn(connect *ProtoConnect) {
	connect.Parser = conn.Parser
	conn.gameConn = connect
}

func (conn *ClientConn) SetGameID(id common.GameServerID) {
	conn.GameID = id
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