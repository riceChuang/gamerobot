package game

import (
	"github.com/riceChuang/gamerobot/common"
	src2 "github.com/riceChuang/gamerobot/service/game/src"
	"sync"
)

type createGameLogicFunc func(gType common.GameServerID) (instance src2.GameLogicBase)

var (
	gameLogicInstanceMap = map[common.GameServerID]src2.GameLogicBase{}
	mu                   = sync.Mutex{}
)

func InitGameInstance() {
	for gameID, _ := range common.GameServerID_Name {
		_ = GetInstanceByContentType(gameID)
	}
	return
}

func GetInstanceByContentType(gameID common.GameServerID) (instance src2.GameLogicBase) {
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

func GetAllGameInstance() map[common.GameServerID]src2.GameLogicBase {
	return gameLogicInstanceMap
}

func createGameRewardInstance(gType common.GameServerID) (instance src2.GameLogicBase) {
	var f createGameLogicFunc
	f = src2.CreateInstanceByContentType
	return f(gType)
}
