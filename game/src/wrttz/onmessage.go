package wrttz

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util"
)

func (wrttz *WRTTZLogic) ParseState(data interface{}) {
	msgFrom := data.(*model.WSMessage)
	msgFrom.From = common.Game
	msgFrom.To = common.GameSender
	state := msgFrom.Msg.Data.(*netproto.WRTTZ_State)
	msgTo := &model.WSMessage{}

	if wrttz.robot == nil{
		wrttz.robot = GetRobtInstanceByType(RobotAllin,wrttz.RoomIndex)
	}

	switch state.GetState() {
	case netproto.WRTTZ_Flow_StartBet:
		{
			betArr := wrttz.robot.GetBetArr()
			bId := int32(netproto.MessageBClassID_Game)
			sId := int32(netproto.WRTTZ_GameMessageClassID_WRTTZBetID)
			info := &model.Message{
				BClassID: bId,
				SClassID: sId,
				Data:     betArr,
			}

			msgTo.From = common.Game
			msgTo.To = common.GameSender
			msgTo.ClientID = msgFrom.ClientID
			msgTo.Msg = info

			util.GetGameDispatcher().AddMessage(msgTo)
		}
	default:
		{

		}
	}
}

func (wrttz *WRTTZLogic) GetBetArr() *netproto.WRTTZ_BetArr {
	return wrttz.robot.GetBetArr()
}

//根據策略取的下注
func (wrttz *WRTTZLogic) ParseGameResult(data interface{}) {
	msgFrom := data.(*model.WSMessage)
	msgFrom.From = common.Game
	msgFrom.To = common.GameSender
	result := msgFrom.Msg.Data.(*netproto.WRTTZ_GameResult)

	//給robot處理
	wrttz.robot.ParseJiang(result.Result)
}


func (wrttz *WRTTZLogic) ParseResultList(data interface{}){
	msgFrom := data.(*model.WSMessage)
	msgFrom.From = common.Game
	msgFrom.To = common.GameSender
	result := msgFrom.Msg.Data.(*netproto.WRTTZ_ResultList)
	wrttz.robot.ParseResultList(result.ResultList)
}