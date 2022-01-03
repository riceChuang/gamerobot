package src

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

type BYDHLogic struct {
	*GameBase
}

func NewBYDHLogic() (instance GameLogicBase) {
	return &BYDHLogic{
		GameBase: NewGameBase(common.GameID_BYDH),
	}
}

func (bydh *BYDHLogic)GetMessageBtn() map[string]*model.Message{
	return map[string]*model.Message{
		netproto.BYDH_GameMessageClassID_name[int32(netproto.BYDH_GameMessageClassID_BYDHRewardZhenZhuID)]: {
			BClassID: int32(netproto.MessageBClassID_GameRoom),
			SClassID: int32(netproto.BYDH_GameMessageClassID_BYDHRewardZhenZhuID),
		},
		netproto.BYDH_GameMessageClassID_name[int32(netproto.BYDH_GameMessageClassID_BYDHBoomRetID)]: {
			BClassID: int32(netproto.MessageBClassID_GameRoom),
			SClassID: int32(netproto.BYDH_GameMessageClassID_BYDHBoomRetID),
		},
		netproto.BYDH_GameMessageClassID_name[int32(netproto.BYDH_GameMessageClassID_BYDHRewardBaoXiangID)]: {
			BClassID: int32(netproto.MessageBClassID_GameRoom),
			SClassID: int32(netproto.BYDH_GameMessageClassID_BYDHRewardBaoXiangID),
		},
	}
}

func (bydh *BYDHLogic) HandleMessage() {

}

func (bydh *BYDHLogic) TransferMessage() {

}