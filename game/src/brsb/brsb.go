package brsb

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util/config"
)

type strategy struct {
	Betodd string `json:"betodd"`
	Params string `json:"params"`
}

type BRSBLogic struct {
	*gamehandler.Base
	Strategy *strategy
}

func NewBRSBLogic(roomIndex int32) gamehandler.GameLogicBase {
	brsb := &BRSBLogic{
		Base:     gamehandler.NewBase(common.GameID_WRTTZ, roomIndex),
		Strategy: &strategy{},
	}
	brsb.Transfer = transfer
	brsb.Init()
	return brsb
}

func transfer(sClass int32) string {
	return netproto.WRTTZ_GameMessageClassID(sClass).String()
}

func (brsb *BRSBLogic) SetStrategy(data interface{}) error {
	if data == nil {
		roomCfg := config.GetGameInstanceByID(brsb.GetGameRoomIndex())
		if roomCfg == nil {
			brsb.Logger.Debug("cant find roomid")
			return nil
		}
		return brsb.UnmarshalConfig(roomCfg.Strategy, brsb.Strategy)
	}

	err := brsb.UnmarshalConfig(data, brsb.Strategy)
	return err
}

func (brsb *BRSBLogic) Init() {
	err := brsb.SetStrategy(nil)
	if err != nil {
		brsb.Logger.Errorf("can't set Strategy game:%v", brsb.GameID)
	}

	brsb.Register(
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBSettingID),
				MsgType:  &netproto.BRSB_Setting{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBStateID),
				MsgType:  &netproto.BRSB_State{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBBetID),
				MsgType:  &netproto.BRSB_Bet{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBBroadCastBetID),
				MsgType:  &netproto.BRSB_BroadCastBetArr{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBGameResultID),
				MsgType:  &netproto.BRSB_GameResult{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBGameResultUserMoneyID),
				MsgType:  &netproto.BRSB_GameResultUserMoney{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBBetAckID),
				MsgType:  &netproto.BRSB_BetAck{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBResultListID),
				MsgType:  &netproto.BRSB_ResultList{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBGameNoInfoID),
				MsgType:  &netproto.BRSB_ResultList{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.BRSB_GameMessageClassID_BRSBSceneID),
				MsgType:  &netproto.BRSB_Scene{},
			},
		},
	)
}
