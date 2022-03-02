package wrttz

import (
	"github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/using/netproto"
	"math/rand"
	"strconv"
	"time"
)

type SingleCornerRobot struct {
	roomIndex     int32
	currentCorners []int32 ///change to corners
	winTime[]       int
	lostTime[]      int
	diff          int
}

func NewSingleCornerRobot(roomIndex int32) *SingleCornerRobot {
	rand.Seed(time.Now().UnixNano())

	robot := &SingleCornerRobot{}
	//隨機選
	robot.roomIndex = roomIndex
	arr := make([]int32,0)
	for i := 0;i<len(BetTypeCorner);i++{
		arr = append(arr,BetTypeCorner[i])
	}
	robot.currentCorners = arr
	robot.winTime = make([]int,3)
	robot.lostTime = make([]int,3)

	return robot
}

// "2"
func (rr *SingleCornerRobot) ParseParam(params string) {
	diff, err := strconv.Atoi(params)
	if err != nil {
		rr.diff = DefaultDiff
	} else {
		rr.diff = diff
	}
}

func (rr *SingleCornerRobot) GetBetArr() *netproto.WRTTZ_BetArr {
	for i := 0;i<len(rr.currentCorners);i++{
		if (rr.winTime[i] - rr.lostTime[i]) < rr.diff {
			continue
		} else if (rr.winTime[i] - rr.lostTime[i]) == rr.diff {
			newCorner := rr.changeBetCorner(rr.currentCorners[i])
			rr.winTime[i] = 0
			rr.lostTime[i] = 0
			rr.currentCorners[i] = newCorner
		}
	}
	index := rand.Intn(len(rr.currentCorners))
	betArr := make([]*netproto.WRTTZ_Bet, 0)
	betRangeArr := BetRange[rr.roomIndex]
	betAmount := betRangeArr[rand.Intn(len(betRangeArr))]
	var betCorner int32

	betCorner = rr.currentCorners[index]

	betArr = append(betArr, &netproto.WRTTZ_Bet{
		BetType: Int32ToBetType(betCorner),
		Score:   proto.Int64(betAmount),
	})
	return &netproto.WRTTZ_BetArr{BetArr: betArr}
}

func (rr *SingleCornerRobot) ParseJiang(arr []*netproto.WRTTZ_BetTypeJiang) {
	for i := 0; i < len(arr); i++ {
		for j := 0;j<len(rr.currentCorners);j++{
			if arr[i].GetBetType() == *Int32ToBetType(rr.currentCorners[j]) {
				if arr[i].GetJiang() == netproto.WRTTZ_JiangType_JiangWin {
					rr.winTime[j]++
				}
				if arr[i].GetJiang() == netproto.WRTTZ_JiangType_JiangLose {
					rr.lostTime[j]++
				}
			}
		}
	}
}

func (rr *SingleCornerRobot)ParseResultList(aee []*netproto.WRTTZ_GameResult){

}

//更換
func (rr *SingleCornerRobot) changeBetCorner(betCorner int32) int32 {
	switch betCorner {
	case BeTypeShangJiao:
		return BeTypeXiaJiao
	case BeTypeXiaJiao:
		return BeTypeGiao
	case BeTypeGiao:
		return BeTypeShangJiao
	default:
		return BeTypeShangJiao
	}
}

type MultiCornerRobot struct {
	roomIndex   int32
	totalAmount int64
}

func NewMultiCornerRobot(roomIndex int32) *MultiCornerRobot {
	rand.Seed(time.Now().UnixNano())

	robot := &MultiCornerRobot{}
	//隨機選
	robot.roomIndex = roomIndex

	return robot
}
//2000
func (rr *MultiCornerRobot) ParseParam(params string) {
	total, err := strconv.Atoi(params)
	if err != nil {
		betRange := BetRange[rr.roomIndex]
		rr.totalAmount = betRange[len(betRange)]
	} else {
		rr.totalAmount = int64(total)
	}
}

func (rr *MultiCornerRobot) ParseJiang(arr []*netproto.WRTTZ_BetTypeJiang) {

}

func (rr *MultiCornerRobot) GetBetArr() *netproto.WRTTZ_BetArr {
	betArr := make([]*netproto.WRTTZ_Bet, 0)

	//決定開幾門
	nums := rand.Intn(len(BetTypeCorner)+1)
	if nums == 0 {
		nums = 1
	}

	betMap := make(map[int32]bool)

	for i := 0; i < len(BetTypeCorner); i++ {
		betMap[BetTypeCorner[i]] = true
	}

	bets := BetStg[rr.roomIndex][nums]
	betCombine := bets[rand.Intn(len(bets))]

	for i := 0;i<len(betCombine);i++{
		betType := MapRandomKeyGet(betMap).(int32)
		delete(betMap,betType)
		betArr = append(betArr, &netproto.WRTTZ_Bet{
			BetType: Int32ToBetType(betType),
			Score:   &betCombine[i],
		})
	}

	return &netproto.WRTTZ_BetArr{BetArr: betArr}
}


func (rr *MultiCornerRobot)ParseResultList(aee []*netproto.WRTTZ_GameResult){

}
