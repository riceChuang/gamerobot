package gamefactory

import "github.com/riceChuang/gamerobot/common"

// Base of stragedy
type Base struct {
	Gameid common.GameServerID
}

// NewBase ...
func NewBase(gameID common.GameServerID) *Base {
	return &Base{
		Gameid: gameID,
	}
}


func (b *Base) GameID() int32 {
	return int32(b.Gameid)
}

func (b *Base) GameName() string {
	return common.GameServerIDToString(b.Gameid)
}


