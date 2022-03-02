package gametype

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/game/src/brsb"
	"github.com/riceChuang/gamerobot/game/src/bydh"
	"github.com/riceChuang/gamerobot/game/src/qzpn"
	"github.com/riceChuang/gamerobot/game/src/wrttz"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/config"
)

func GetInstanceByContentType(gType common.GameServerID, gIndex ...int32) (instance GameType) {

	var cfg = model.CommonCfg{}
	gameList := config.GetInstance().GameList
	for _, config := range gameList {
		if config.GameID == int32(gType) {
			cfg = config
		}
	}

	var index int32
	if len(gIndex) > 0 {
		index = gIndex[0]
	}

	switch gType {
	case common.GameID_QZPN:
		return NewFightBase(cfg, qzpn.NewQZPNLogic(index))
	case common.GameID_BYDH:
		return NewFightBase(cfg, bydh.NewBYDHLogic(index))
	case common.GameID_WRTTZ:
		return NewBaiRenBase(cfg, wrttz.NewWRTTZLogic(index))
	case common.GameID_BRSB:
		return NewBaiRenBase(cfg, brsb.NewBRSBLogic(index))
	default:
		return nil
	}
	return nil
}
