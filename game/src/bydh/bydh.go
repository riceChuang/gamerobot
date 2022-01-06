package bydh

import (
	"github.com/riceChuang/gamerobot/framework/gamefactory"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

type BYDHLogic struct {
	*gamefactory.Base
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

func (bydh *BYDHLogic) HandleMessage(message *model.WSMessage) {

}
