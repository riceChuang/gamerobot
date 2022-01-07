package gametype

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/connection/connecttype"
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sort"
)

type GameType interface {
	GetGameConfig() model.CommonCfg                                 //取得遊戲config
	Initialize(gameConnect *connecttype.GameConnect)                //初始化遊戲內部註冊
	HandleMessage(message *model.WSMessage)                         //處理各自遊戲訊息
	UnmarshalMessage(message *model.Message) (proto.Message, error) //解析封包
	GetGameHandler() gamehandler.GameLogicBase                      //取得遊戲handler
}

// Base of stragedy
type BaseGameType struct {
	gid         common.GameServerID
	gName       string
	gRoomIndex  map[int]string
	gHallPort   int32
	gameHandler gamehandler.GameLogicBase
	gameConnect *connecttype.GameConnect
	logger      *log.Entry
}

// NewBase ...
func NewBaseGameType(cfg model.CommonCfg, gamelogic gamehandler.GameLogicBase) *BaseGameType {
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

//Initialize.....
func (bt *BaseGameType) Initialize(gameConnect *connecttype.GameConnect) {
	bt.gameConnect = gameConnect
	return
}

//get game config
func (bt *BaseGameType) GetGameConfig() model.CommonCfg {
	gcfg := model.CommonCfg{
		GameID:   int32(bt.gid),
		GameName: bt.gName,
		HallPort: bt.gHallPort,
		RoomType: make([]string, 0, len(bt.gRoomIndex)),
	}

	keys := make([]int, 0, len(bt.gRoomIndex))
	for k := range bt.gRoomIndex {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		gcfg.RoomType = append(gcfg.RoomType, bt.gRoomIndex[k])
	}
	return gcfg
}

// HandleMessage by Websocket or client
func (bt *BaseGameType) HandleMessage(message *model.WSMessage) {

}

func (bt *BaseGameType) GetGameHandler() gamehandler.GameLogicBase {
	return bt.gameHandler
}

func (bt *BaseGameType) UnmarshalMessage(message *model.Message) (proto.Message, error) {
	return bt.gameHandler.UnmarshalMessage(message)
}
