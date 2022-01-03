package src

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/config"
	"sort"
)

type GameLogicBase interface {
	GetGameConfig() model.CommonCfg
	GetMessageBtn() map[string]*model.Message
	HandleMessage()
	TransferMessage()
}

func CreateInstanceByContentType(gType common.GameServerID) (instance GameLogicBase) {
	cfg := config.GetInstance()
	var gamecfg = &model.CommonCfg{}
	for _, game := range cfg.GameList {
		if gType == common.GameServerID(game.GameID) {
			gamecfg = &game
			break
		}
	}
	switch gType {
	case common.GameID_QZPN:
		return NewQZPNLogic(gamecfg)
	case common.GameID_BYDH:
		return NewBYDHLogic(gamecfg)
	default:
		return nil
	}
}

type GameBase struct {
	gid        common.GameServerID
	gName      string
	gRoomIndex map[int]string
	gHallPort  int32
}

func NewGameBase(cfg *model.CommonCfg) *GameBase {
	gb := &GameBase{
		gid:        common.GameServerID(cfg.GameID),
		gName:      common.GameServerIDToString(common.GameServerID(cfg.GameID)),
		gHallPort:  cfg.HallPort,
		gRoomIndex: make(map[int]string),
	}
	for index, name := range cfg.RoomType {
		gb.gRoomIndex[index] = name
	}
	return gb
}

func (gb *GameBase) GetGameConfig() model.CommonCfg {
	gcfg := model.CommonCfg{
		GameID:   int32(gb.gid),
		GameName: gb.gName,
		HallPort: gb.gHallPort,
		RoomType: make([]string, 0, len(gb.gRoomIndex)),
	}

	keys := make([]int, 0, len(gb.gRoomIndex))
	for k := range gb.gRoomIndex {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		gcfg.RoomType = append(gcfg.RoomType, gb.gRoomIndex[k])
	}
	return gcfg
}

func (gb *GameBase) HandleMessage() {

}
