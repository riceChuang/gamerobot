package wrttz

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util/config"
	"strconv"
)

type strategy struct {
	Betodd string `json:"betodd"`
	Params string  `json:"params"`
}

type WRTTZLogic struct {
	*gamehandler.Base
	Strategy *strategy
	robot    Robot
}

func NewWRTTZLogic(roomIndex int32) gamehandler.GameLogicBase {
	wrttz := &WRTTZLogic{
		Base:     gamehandler.NewBase(common.GameID_WRTTZ, roomIndex),
		Strategy: &strategy{},
	}
	wrttz.Transfer = transfer
	wrttz.Init()
	return wrttz
}

func transfer(sClass int32) string {
	return netproto.WRTTZ_GameMessageClassID(sClass).String()
}

func (wrttz *WRTTZLogic) SetStrategy(data interface{}) error {
	if data == nil {
		roomCfg := config.GetGameInstanceByID(wrttz.GetGameRoomIndex())
		if roomCfg == nil {
			wrttz.Logger.Debug("cant find roomid")
			return nil
		}
		return wrttz.UnmarshalConfig(roomCfg.Strategy, wrttz.Strategy)
	}

	err := wrttz.UnmarshalConfig(data, wrttz.Strategy)
	robotType,_:= strconv.Atoi(wrttz.Strategy.Betodd)
	//根據下注策略件機器人
	wrttz.robot = GetRobtInstanceByType(robotType,wrttz.RoomIndex)
	wrttz.robot.ParseParam(wrttz.Strategy.Params)
	return err
}

func (wrttz *WRTTZLogic) Init() {
	err := wrttz.SetStrategy(nil)
	if err != nil {
		wrttz.Logger.Errorf("can't set Strategy game:%v", wrttz.GameID)
	}
	wrttz.Register(
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZSettingID),
				MsgType:  &netproto.WRTTZ_Setting{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.WRTTZ_GameMessageClassID_WRTTZStateID),
				MsgType:   &netproto.WRTTZ_State{},
				OnMessage: wrttz.ParseState,
			},
			//IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZBetID),
				MsgType:  &netproto.WRTTZ_Bet{},
			},
			//IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZBroadCastBetID),
				MsgType:  &netproto.WRTTZ_BroadCastBetArr{},
			},
			//IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZRtpInfoID),
				MsgType:  &netproto.WRTTZ_RtpInfo{},
			},
		},

		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZGameResultID),
				MsgType:  &netproto.WRTTZ_GameResult{},
				OnMessage: wrttz.ParseGameResult,
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZCardResultId),
				MsgType:  &netproto.WRTTZ_CardResult{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZGameResultUserMoneyID),
				MsgType:  &netproto.WRTTZ_GameResultUserMoney{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZBetAckID),
				MsgType:  &netproto.WRTTZ_BetAck{},
			},
			//IsButton: true,
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZResultListID),
				MsgType:  &netproto.WRTTZ_ResultList{},
				OnMessage: wrttz.ParseResultList,
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID:  int32(netproto.MessageBClassID_Game),
				SClassID:  int32(netproto.WRTTZ_GameMessageClassID_WRTTZGameNoInfoID),
				MsgType:   &netproto.WRTTZ_GameNoInfo{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZSceneID),
				MsgType:  &netproto.WRTTZ_Scene{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.WRTTZ_GameMessageClassID_WRTTZRichestListID),
				MsgType:  &netproto.WRTTZ_RichestList{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.GameMessageClassID_UserStateChangeID),
				MsgType:  &netproto.UserStateChange{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.GameMessageClassID_UserEnterID),
				MsgType:  &netproto.UserEnter{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.GameMessageClassID_UserLeaveID),
				MsgType:  &netproto.UserLeave{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.GameMessageClassID_LookOnUserNotifyID),
				MsgType:  &netproto.LookOnUserNotify{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.GameMessageClassID_GameStartID),

			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.GameMessageClassID_GameEndID),
				MsgType:  &netproto.TestGameEnd{},
			},
		},
		&model.GameMessageInfo{
			Handler: model.Handler{
				BClassID: int32(netproto.MessageBClassID_Game),
				SClassID: int32(netproto.GameMessageClassID_GameStartID),
				MsgType:  &netproto.GameServerStart{},
			},
		},
	)
}
