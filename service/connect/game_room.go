package connect

import (
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type GameConnectInterface interface {
}

type GameConnect struct {
	logger   *log.Entry
	GameWS   *framework.ProtoConnect
	mu       sync.Mutex
	userInfo *netproto.UserLoginRet
}

func NewGameConnect(gameWS *framework.ProtoConnect, uinfo *netproto.UserLoginRet) *GameConnect {
	return &GameConnect{
		GameWS: gameWS,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "GameConnect",
		}),
		userInfo: uinfo,
	}
}

// InitGameWs ...
func (gc *GameConnect) ConnectGameWs() {
	gc.GameWS.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_GameRoom),
		SClassID:  int32(netproto.GameRoomClassID_LoginRoomRetID),
		MsgType:   &netproto.LoginGameRoomRet{},
		OnMessage: gc.onGameLogin,
	})

	gc.GameWS.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_GameRoom),
		OnMessage: gc.onGameMessage,
	})

	gc.GameWS.Connect()
	go func() {
		time.Sleep(2 * time.Second)
		gc.requestLoginGame()
	}()

}

//request GameWs
func (gc *GameConnect) requestLoginGame() {
	gc.logger.Infof("[AAAAAAAAAAAArequestLoginGame] 請求登入遊戲")
	gc.Request(
		gc.GameWS,
		int32(netproto.MessageBClassID_GameRoom),
		int32(netproto.GameRoomClassID_LoginRoomID),
		&netproto.LoginGameRoomInfo{
			UserID: gc.userInfo.UserID,
			Cer:    gc.userInfo.Cer,
			HDCode: gc.userInfo.HDCode,
			HDType: gc.userInfo.HDType,
		},
	)
}

func (gc *GameConnect) onGameLogin(msg interface{}) {
	data, ok := msg.(*netproto.LoginGameRoomRet)
	if !ok {
		log.Println("Error: Data Type Wrong!")
		return
	}
	gc.logger.Infof("[requestLoginGame收到遊戲回覆]onGameLogin:%v, user:%v", data.String(), gc.userInfo.String())
}

func (gc *GameConnect) onGameMessage(msg interface{}) {
	message, ok := msg.(*model.Message)
	if !ok {
		gc.logger.Error("Error: data transfer fail")
	}
	gc.logger.Infof("收到game消息封包 data:%v", message.GetName())
}

// Request send and handle protoMsg Error
func (gc *GameConnect) Request(ws *framework.ProtoConnect, bclsID int32, sclsID int32, protoData interface{}) {
	message, err := model.NewProto(bclsID, sclsID, protoData)

	// should kill robot
	if err != nil {
		gc.logger.Panic(err.Error())
	}
	ws.Write(message)
}

// CleanGs ...
func (gc *GameConnect) CleanGs() {
	gc.logger.Info("[DEBUG][Manager][CleanGs] ----")
	if gc.GameWS != nil {
		gc.GameWS.Clean()
		gc.GameWS = nil
	}
}
