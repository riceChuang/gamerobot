package src

import (
	"gitlab.baifu-tech.net/dsg-game/game-robot/common"
	"gitlab.baifu-tech.net/dsg-game/game-robot/model"
)

type GameLogicBase interface {
	GameID() int32
	GameName() string
	GetMessageBtn() map[string]*model.Message
	HandleMessage()
	TransferMessage()
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


type GameBase struct {
	gid common.GameServerID
}

func NewGameBase(gameID common.GameServerID) *GameBase {
	return &GameBase{
		gid: gameID,
	}
}


func (gb *GameBase) GameID() int32 {
	return int32(gb.gid)
}

func (gb *GameBase) GameName() string {
	return common.GameServerIDToString(gb.gid)
}


func (gb *GameBase) HandleMessage() {

}