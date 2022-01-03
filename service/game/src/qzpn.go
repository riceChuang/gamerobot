package src

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

type QZPNLogic struct {
	*GameBase
}

func NewQZPNLogic() (instance GameLogicBase) {
	return &QZPNLogic{
		GameBase: NewGameBase(common.GameID_QZPN),
	}
}



func (qzpn *QZPNLogic) GetMessageBtn() map[string]*model.Message {
	return map[string]*model.Message{
		netproto.QZPN_GameMessageClassID_name[int32(netproto.QZPN_GameMessageClassID_QZPNBroadCastGameEndResultID)]: {
			BClassID: int32(netproto.MessageBClassID_GameRoom),
			SClassID: int32(netproto.QZPN_GameMessageClassID_QZPNBroadCastGameEndResultID),
		},
		netproto.QZPN_GameMessageClassID_name[int32(netproto.QZPN_GameMessageClassID_QZPNGameResultID)]: {
			BClassID: int32(netproto.MessageBClassID_GameRoom),
			SClassID: int32(netproto.QZPN_GameMessageClassID_QZPNGameResultID),
		},
		netproto.QZPN_GameMessageClassID_name[int32(netproto.QZPN_GameMessageClassID_QZPNXianjiaBetID)]: {
			BClassID: int32(netproto.MessageBClassID_GameRoom),
			SClassID: int32(netproto.QZPN_GameMessageClassID_QZPNXianjiaBetID),
		},
	}
}


func (qzpn *QZPNLogic) TransferMessage() {

}