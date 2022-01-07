package qzpn

import (
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

type QZPNLogic struct {
	*gamehandler.Base
}

func NewQZPNLogic() gamehandler.GameLogicBase {
	qzpn := &QZPNLogic{
		Base: gamehandler.NewBase("QZPN", transfer),
	}
	qzpn.Init()
	return qzpn
}

func transfer(sClass int32) string {
	return netproto.QZPN_GameMessageClassID(sClass).String()
}

func (qzpn *QZPNLogic) Init() {
	qzpn.Register(
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNSettingID),
				MsgType:   &netproto.QZPN_GameSetting{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNStateID),
				MsgType:   &netproto.QZPN_Sate{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNQiangZhuangID),
				MsgType:   &netproto.QZPN_QiangZhuang{},
				OnMessage: qzpn.hiThree,
			},
			IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNOnUserZhuangBetID),
				MsgType:   &netproto.QZPN_UserBet{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNBroadCastZhuangID),
				MsgType:   &netproto.QZPN_BroadCastQiangZhuang{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNXianjiaBetID),
				MsgType:   &netproto.QZPN_XianJiaBetTimes{},
			},
			IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNOnUserXianBetID),
				MsgType:   &netproto.QZPN_UserBet{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNBroadCastXianBetID),
				MsgType:   &netproto.QZPN_BroadCastXianBetTimes{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNBroadCastDiceID),
				MsgType:   &netproto.QZPN_DiceInfo{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNGameResultID),
				MsgType:   &netproto.QZPN_GameResult{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNBroadCastGameEndResultID),
				MsgType:   &netproto.QZPN_GameEndResult{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNGameSceneID),
				MsgType:   &netproto.QZPN_GameScene{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNUserBetInfoID),
				MsgType:   &netproto.QZPN_UserBetInfo{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.QZPN_GameMessageClassID_QZPNGameNo),
				MsgType:   &netproto.QZPN_GameNoInfo{},
			},
		},
	)
}
