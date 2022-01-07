package connecttype

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/connection/connect"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sync"
)

type GameConnectInterface interface {
}

type GameConnect struct {
	ClientID  string
	logger    *log.Entry
	GameWS    *connect.ProtoConnect
	mu        sync.Mutex
	userInfo  *netproto.UserLoginRet
	UserMoney int32
}

func NewGameConnect(gameWS *connect.ProtoConnect, uinfo *netproto.UserLoginRet) *GameConnect {
	gc := &GameConnect{
		GameWS: gameWS,
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "Game_Type_Connect",
		}),
		userInfo: uinfo,
	}
	if uinfo != nil {
		gc.UserMoney = int32(uinfo.GetUserData().GetCashAmount() / int64(100))
	}
	return gc
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
		BClassID:  int32(netproto.MessageBClassID_Game),
		SClassID:  0,
		OnMessage: gc.onGameMessage,
	})

	gc.GameWS.Connect()
	gc.requestLoginGame()
}

// Request send and handle protoMsg Error
func (gc *GameConnect) Request(ws *connect.ProtoConnect, bclsID int32, sclsID int32, protoData interface{}) {
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

func (gc *GameConnect) GetUserID() int32 {
	if gc.userInfo != nil {
		return gc.userInfo.GetUserID()
	}
	return 0
}

//request GameWs
func (gc *GameConnect) requestLoginGame() {
	gc.logger.Infof("[requestLoginGame] 請求登入遊戲")
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
	wsMessage := &model.WSMessage{
		From:     common.Game,
		To:       common.GameServerTransfer,
		ClientID: gc.ClientID,
		Msg:      message,
	}
	util.GetGameDispatcher().AddMessage(wsMessage)
}
