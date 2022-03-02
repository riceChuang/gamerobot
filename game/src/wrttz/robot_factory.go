package wrttz

import (
	"github.com/riceChuang/gamerobot/using/netproto"
)

type Robot interface {
	ParseParam(params string)
	GetBetArr() *netproto.WRTTZ_BetArr
	ParseJiang(arr []*netproto.WRTTZ_BetTypeJiang)
	ParseResultList(aee []*netproto.WRTTZ_GameResult)
}

func GetRobtInstanceByType(robotType int, roomIndex int32) Robot {
	switch robotType {
	case RobotRand:
		{
			return NewRandRobot(roomIndex)
		}
	case RobotAllin:
		{
			return NewAllInRobot(roomIndex)
		}
	case RobotSingleCorner:
		{
			return NewSingleCornerRobot(roomIndex)
		}
	case RobotMultiCorner:
		{
			return NewMultiCornerRobot(roomIndex)
		}
	case RobotDiagonal:
		{
			return NewDiagonalRobot(roomIndex)
		}
	case RobotReverse:
		{
			return NewReverseRobot(roomIndex)
		}
	case RobotDirect:
		{
			return NewDirectRobot(roomIndex)
		}
		//default:
	}
	return NewRandRobot(roomIndex)
}

type RobotStrategy struct {
	RobotType    string
	MustBets     []*RobotBetRange
	PossibleBets []*RobotBetRange
}

type RobotBetRange struct {
	BetType          int
	ModulationAmount int64
}
