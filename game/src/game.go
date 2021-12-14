package src

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
)

type GameLogicBase interface {
	GameID() int32
	GameName() string
	GetMessageBtn() map[string]*model.Message
}

func CreateInstanceByContentType(gType common.GameServerID) (instance GameLogicBase) {
	switch gType {
	case common.GameID_QZPN:
		return NewQZPNLogic()
	case common.GameID_BYDH:
		return NewBYDHLogic()
	default:
		return nil
	}
}
