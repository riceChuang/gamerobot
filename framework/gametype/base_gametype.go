package gametype

import (
	"fmt"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sort"
)

// Base of stragedy
type BaseGameType struct {
	gid           common.GameServerID
	gName         string
	gRoomIndex    map[int]string
	gHallPort     int32
	onGameMessage func(msg *model.WSMessage)
	gameHandler   GameLogicBase
	logger        *log.Entry
}

// NewBase ...
func NewBaseGameType(cfg model.CommonCfg, gamelogic GameLogicBase) *BaseGameType {
	bgt := &BaseGameType{
		gid: common.GameServerID(cfg.GameID),
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": fmt.Sprintf("game-%d-instence", common.GameServerID(cfg.GameID)),
		}),
		gameHandler: gamelogic,
		gName:       common.GameServerIDToString(common.GameServerID(cfg.GameID)),
		gHallPort:   cfg.HallPort,
		gRoomIndex:  make(map[int]string),
	}

	for index, name := range cfg.RoomType {
		bgt.gRoomIndex[index] = name
	}
	return bgt
}

func (fb *BaseGameType) Initialize(ws *framework.ProtoConnect) {

}

func (fb *BaseGameType) GetGameConfig() model.CommonCfg {
	gcfg := model.CommonCfg{
		GameID:   int32(fb.gid),
		GameName: fb.gName,
		HallPort: fb.gHallPort,
		RoomType: make([]string, 0, len(fb.gRoomIndex)),
	}

	keys := make([]int, 0, len(fb.gRoomIndex))
	for k := range fb.gRoomIndex {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		gcfg.RoomType = append(gcfg.RoomType, fb.gRoomIndex[k])
	}
	return gcfg
}

func (fb *BaseGameType) HandleMessage(message *model.WSMessage) {

}

func (fb *BaseGameType) GetGameHandler() GameLogicBase {
	return fb.gameHandler
}

