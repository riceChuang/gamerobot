package game

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/game/src"
	"sync"
)

type createGameLogicFunc func(gType common.GameServerID) (instance src.GameLogicBase)

var (
	gameLogicInstanceMap = map[common.GameServerID]src.GameLogicBase{}
	mu                   = sync.Mutex{}
)

func InitGameInstance() {
	for gameID, _ := range common.GameServerID_Name {
		_ = GetInstanceByContentType(gameID)
	}
	return
}

func GetInstanceByContentType(gameID common.GameServerID) (instance src.GameLogicBase) {
	mu.Lock()
	defer mu.Unlock()
	var (
		exist = false
	)
	if instance, exist = gameLogicInstanceMap[gameID]; !exist {
		instance = createGameRewardInstance(gameID)
		if instance != nil {
			gameLogicInstanceMap[gameID] = instance
		}
	}
	return
}

func GetAllGameInstance() map[common.GameServerID]src.GameLogicBase {
	return gameLogicInstanceMap
}

func createGameRewardInstance(gType common.GameServerID) (instance src.GameLogicBase) {
	var f createGameLogicFunc
	f = src.CreateInstanceByContentType
	return f(gType)
}
