package src

import (
	"errors"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util/config"
)

type strategy struct {
}

type DefaultLogic struct {
	*gamehandler.Base
	Strategy *strategy
}

func NewLogic(roomIndex int32) gamehandler.GameLogicBase {
	d := &DefaultLogic{
		Base: gamehandler.NewBase(common.GameID_QZPN, roomIndex),
	}
	d.Transfer = transfer
	d.Init()
	return d
}

func transfer(sClass int32) string {
	return netproto.GameMessageClassID(sClass).String()
}

func (d *DefaultLogic) SetStrategy(data interface{}) error {
	if data == nil {
		roomCfg := config.GetGameInstanceByID(d.GetGameRoomIndex())
		if roomCfg == nil {
			return errors.New("cant find roomid")
		}
		return d.UnmarshalConfig(roomCfg.Strategy, d.Strategy)
	}
	return d.UnmarshalConfig(data, d.Strategy)
}

func (d *DefaultLogic) Init() {

	err := d.SetStrategy(nil)
	if err != nil {
		d.Logger.Errorf("can't set Strategy game:%v", d.GameID)
	}

	d.Register()
}
