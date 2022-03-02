package bydh

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

type BYDHLogic struct {
	*gamehandler.Base
}

func NewBYDHLogic(roomIndex int32) gamehandler.GameLogicBase {
	bydh := &BYDHLogic{
		Base: gamehandler.NewBase(common.GameID_BYDH, roomIndex),
	}
	bydh.Transfer = transfer
	bydh.Init()
	return bydh
}

func transfer(sClass int32) string {
	return netproto.BYDH_GameMessageClassID(sClass).String()
}

func (bydh *BYDHLogic) Init() {
	bydh.Register(
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BYDH_GameMessageClassID_BYDHRewardZhenZhuID),
				MsgType:  &netproto.UserStateChange{},
			},
			IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BYDH_GameMessageClassID_BYDHBoomRetID),
				MsgType:  &netproto.UserStateChange{},
			},
			IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BYDH_GameMessageClassID_BYDHRewardBaoXiangID),
				MsgType:  &netproto.UserStateChange{},
			},
			IsButton: true,
		},
	)
}
