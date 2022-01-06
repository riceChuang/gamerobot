package gametype

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/game/src/bydh"
	"github.com/riceChuang/gamerobot/game/src/qzpn"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/config"
)

type GameLogicBase interface {
	GetMessageBtn() map[string]*model.Message
	HandleMessage(message *model.WSMessage)
}

func GetInstanceByContentType(gType common.GameServerID) (instance GameType) {

	var cfg = model.CommonCfg{}
	gameList := config.GetInstance().GameList
	for _, config := range gameList {
		if config.GameID == int32(gType) {
			cfg = config
		}
	}

	switch gType {
	case common.GameID_QZPN:
		return NewFightBase(cfg, &qzpn.QZPNLogic{})
	case common.GameID_BYDH:
		return NewFightBase(cfg, &bydh.BYDHLogic{})
	default:
		return nil
	}
	return nil
}
