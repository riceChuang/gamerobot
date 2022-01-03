package connect

import (
	"context"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type ClientConn struct {
	ctx       context.Context
	ID        string
	wsConn    *HttpConnect
	gameConn  *ProtoConnect
	Parser    framework.Parser
	GameID    common.GameServerID
	logger    *log.Entry
	lock      sync.Mutex // readwrite lock
	timerChan chan bool
}

func NewClientWsGameConn(ctx context.Context) (*ClientConn, error) {
	client := &ClientConn{
		ID:  uuid.NewV4().String(),
		ctx: ctx,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "client_ws",
		}),
		timerChan: make(chan bool, 1),
	}
	return client, nil
}

func (client *ClientConn) SetWsConn(connect *HttpConnect) {
	connect.ClientID = client.ID
	client.wsConn = connect
}

func (client *ClientConn) SetGameConn(connect *ProtoConnect) {
	connect.Parser = client.Parser
	client.gameConn = connect
}

func (client *ClientConn) SetGameID(id common.GameServerID) {
	client.GameID = id
}

func (client *ClientConn) ProtoConnect() error {
	err := client.gameConn.Connect()
	if err != nil {
		return err
	}
	client.initHeartBeat()
	return nil
}

func (client *ClientConn) initHeartBeat() {
	client.gameConn.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_Common),
		SClassID:  int32(netproto.PlatformCommonClassID_HeartBeatConfigID),
		MsgType:   &netproto.HeartBeatConfig{},
		OnMessage: client.onHeartBeatCfg,
	})

	client.requestHeartBeatCfg()
}

func (client *ClientConn) requestHeartBeatCfg() {
	message, err := model.NewProto(
		int32(netproto.MessageBClassID_Common),
		int32(netproto.PlatformCommonClassID_RequestHeartBeatConfigID),
		nil,
	)

	if err != nil {
		client.gameConn.logger.Println("Error: send heart beat failed", err.Error())
	}
	client.gameConn.Write(message)
}

func (client *ClientConn) onHeartBeatCfg(b interface{}) {
	data, ok := b.(*netproto.HeartBeatConfig)
	if !ok {
		// should kill robot
		client.gameConn.logger.Panic("Error: Data Type Wrong!")
		return
	}
	client.gameConn.logger.Println("onHeartBeatCfg:", data.String())

	if client.timerChan != nil {
		client.timerChan <- true
	}
	client.timerChan = util.NewTicker(client.sendHeartBeat, time.Duration(data.GetBeatInterval())*time.Second)
}

func (client *ClientConn) sendHeartBeat() {
	message, err := model.NewProto(
		int32(netproto.MessageBClassID_Common),
		int32(netproto.PlatformCommonClassID_HeartBeatID),
		nil,
	)

	if err != nil {
		client.gameConn.logger.Println("Error: send heart beat failed", err.Error())
	}
	client.gameConn.Write(message)
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
