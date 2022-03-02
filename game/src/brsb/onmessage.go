package brsb

import (
	"github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util"
)

func (brsb *BRSBLogic) ParseState(data interface{}) {
	msgFrom := data.(*model.WSMessage)
	msgFrom.From = common.Game
	msgFrom.To = common.GameSender
	state := msgFrom.Msg.Data.(*netproto.BRSB_State)
	msgTo := &model.WSMessage{}

	switch state.GetState() {
	case netproto.BRSB_Flow_BRSBBet:
		{
			betArr := getBetArr()
			bId := int32(netproto.MessageBClassID_Game)
			sId := int32(netproto.BRSB_GameMessageClassID_BRSBBetID)
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

func getBetArr() []*netproto.BRSB_Bet {
	betArr := []*netproto.BRSB_Bet{
		{
			BetType: netproto.BRSB_BetType_BRSB_Bet_Type_Dan.Enum(),
			Score:   proto.Int64(60000),
		},
		{
			BetType: netproto.BRSB_BetType_BRSB_Bet_Type_Shuang.Enum(),
			Score:   proto.Int64(10000),
		},
	}
	return betArr
}
