package gametype

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/connection/connecttype"
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

// Base of stragedy
type FightBase struct {
	*BaseGameType
}

// NewBase ...
func NewFightBase(cfg model.CommonCfg, gamelogic gamehandler.GameLogicBase) GameType {
	fb := &FightBase{
		BaseGameType: NewBaseGameType(cfg, gamelogic),
	}

	return fb
}

func (fb *FightBase) Initialize(gameConnect *connecttype.GameConnect) {
	fb.logger.Info("初始化 fightBase")
	if gameConnect != nil {
		fb.gameConnect = gameConnect
		gameConnect.GameWS.Register(&model.Handler{
			BClassID:  int32(netproto.MessageBClassID_Game),
			SClassID:  int32(netproto.GameMessageClassID_UserStateChangeID),
			MsgType:   &netproto.UserStateChange{},
			OnMessage: fb.onUserStateChange,
		})
	}
}

func (fb *FightBase) HandleMessage(message *model.WSMessage) {

	switch message.Msg.SClassID {
	default:
		fb.gameHandler.HandleMessage(message)
	}
	return
}

func (fb *FightBase) onUserStateChange(msg interface{}) {
	data, ok := msg.(*netproto.UserStateChange)
	if !ok {
		fb.logger.Error("onUserStateChange Error: Data Type Wrong!")
		return
	}

	user := data.GetUser()
	if user.GetUserID() == fb.gameConnect.GetUserID() {
		fb.gameConnect.UserMoney = int32(user.GetScore() / 100)
		fb.logger.Info("onUserStateChange User update money: %d", fb.gameConnect.UserMoney)
		fb.logger.Info("onUserStateChange User current stateValue: %d", user.GetStateValue())
		if user.GetStateValue() == int32(common.StateValueSit) {
			fb.sendReady()
		}
	}
}

func (fb *FightBase) sendReady() {
	if fb.gameConnect != nil {
		fb.gameConnect.Request(
			fb.gameConnect.GameWS,
			int32(netproto.MessageBClassID_GameRoom),
			int32(netproto.GameRoomClassID_GameReadyID),
			nil,
		)
	}
	return
}
