package gametype

import (
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

// Base of stragedy
type FightBase struct {
	*BaseGameType
}

// NewBase ...
func NewFightBase(cfg model.CommonCfg, gamelogic GameLogicBase) GameType {
	fb := &FightBase{
		BaseGameType: NewBaseGameType(cfg,gamelogic),
	}
	return fb
}

func (fb *FightBase) Initialize(ws *framework.ProtoConnect) {
	if ws != nil {
		ws.Register(&model.Handler{
			BClassID:  int32(netproto.MessageBClassID_Game),
			SClassID:  int32(netproto.GameMessageClassID_UserStateChangeID),
			MsgType:   &netproto.UserStateChange{},
			OnMessage: fb.onUserStateChange,
		})
	}
}

func (fb *FightBase) HandleMessage(message *model.WSMessage) {

	switch message.Msg.SClassID {
	case int32(netproto.GameMessageClassID_UserStateChangeID):
		fb.onUserStateChange(message)
	}

}

func (fb *FightBase) onUserStateChange(msg interface{}) {
	//data, ok := msg.(*netproto.UserStateChange)
	//if !ok {
	//	fb.logger.Error("onUserStateChange Error: Data Type Wrong!")
	//	return
	//}

	//user := data.GetUser()
	//if user.GetUserID() == f.NetManager.GetUserID() {
	//f.NetManager.UserMoney = int32(user.GetScore() / 100)
	//f.NetManager.Printf("onUserStateChange Robot update money: %d", f.NetManager.UserMoney)
	//f.NetManager.Printf("onUserStateChange Robot current stateValue: %d", user.GetStateValue())
	//if user.GetStateValue() == StateValueSit {
	//	f.sendReady()
	//}
	//}
}

func (fb *FightBase) sendReady() {

	//fb.NetManager.Request(
	//	ws,
	//	int32(netproto.MessageBClassID_GameRoom),
	//	int32(netproto.GameRoomClassID_GameReadyID),
	//	nil,
	//)
	//wsMsg := model.WSMessage{
	//	From:
	//	To:
	//}
	//

}
