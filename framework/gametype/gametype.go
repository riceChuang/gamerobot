package gametype

import (
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
)

type GameType interface {
	GetGameConfig() model.CommonCfg
	Initialize(ws *framework.ProtoConnect)
	HandleMessage(message *model.WSMessage)
	GetGameHandler() GameLogicBase
}


