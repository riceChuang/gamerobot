package gametype

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/game/src/bydh"
	"github.com/riceChuang/gamerobot/game/src/qzpn"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/config"
)

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
		return NewFightBase(cfg, qzpn.NewQZPNLogic())
	case common.GameID_BYDH:
		return NewFightBase(cfg, bydh.NewBYDHLogic())
	default:
		return nil
	}
	return nil
}
