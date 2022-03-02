package wrttz

import (
	"github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/using/netproto"
	"math/rand"
	"strconv"
	"time"
)

type DiagonalRobot struct {
	roomIndex              int32
	betAmount              int64
	continuesShangMen      int
	continuesBeTypeTianMen int
	continuesBeTypeXiaMen  int
}

func NewDiagonalRobot(roomIndex int32) *DiagonalRobot {
	rand.Seed(time.Now().UnixNano())

	robot := &DiagonalRobot{
		roomIndex: roomIndex,
	}

	return robot
}

// param:"10000"
func (rr *DiagonalRobot) ParseParam(params string) {
	betAmount, err := strconv.Atoi(params)
	if err != nil {
		betRange := BetRange[rr.roomIndex]
		rr.betAmount = betRange[len(betRange)]
	} else {
		rr.betAmount = int64(betAmount)
	}
}

func (rr *DiagonalRobot) ParseJiang([]*netproto.WRTTZ_BetTypeJiang) {

}

func (rr *DiagonalRobot) GetBetArr() *netproto.WRTTZ_BetArr {
	betArr := make([]*netproto.WRTTZ_Bet, 0)
	betNums := decideBetNums(rr.continuesBeTypeXiaMen, rr.continuesShangMen, rr.continuesBeTypeXiaMen)
	switch betNums {
	case 1:
		{
			//單門
			maxNum := findMax(rr.continuesBeTypeXiaMen, rr.continuesBeTypeTianMen, rr.continuesShangMen)

			switch maxNum {
			case rr.continuesBeTypeXiaMen:
				betType1 := rr.changeBetCorner(BeTypeXiaMen)
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType1),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				return &netproto.WRTTZ_BetArr{BetArr: betArr}
			case rr.continuesBeTypeTianMen:
				betType1 := rr.changeBetCorner(BeTypeTianMen)
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType1),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				return &netproto.WRTTZ_BetArr{BetArr: betArr}
			case rr.continuesShangMen:
				betType1 := rr.changeBetCorner(BeTypeShangMen)
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType1),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				return &netproto.WRTTZ_BetArr{BetArr: betArr}
			}
		}
	case 2:
		{
			//兩門
			if compareVal(rr.continuesShangMen, rr.continuesBeTypeTianMen) {
				//下
				betType1 := rr.changeBetCorner(BeTypeShangMen)
				//橋
				betType2 := rr.changeBetCorner(BeTypeTianMen)
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType1),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType2),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				return &netproto.WRTTZ_BetArr{BetArr: betArr}
			}

			if compareVal(rr.continuesShangMen, rr.continuesBeTypeXiaMen) {
				//下
				betType1 := rr.changeBetCorner(BeTypeShangMen)
				//上角
				betType2 := rr.changeBetCorner(BeTypeXiaMen)
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType1),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType2),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				return &netproto.WRTTZ_BetArr{BetArr: betArr}
			}

			if compareVal(rr.continuesBeTypeTianMen, rr.continuesBeTypeXiaMen) {
				//下
				betType1 := rr.changeBetCorner(BeTypeShangMen)
				//上角
				betType2 := rr.changeBetCorner(BeTypeXiaMen)
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType1),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(betType2),
					Score:   proto.Int64(rr.betAmount / 2),
				})
				return &netproto.WRTTZ_BetArr{BetArr: betArr}
			}
		}
	case 3:
		{
			//三門同分
			if rr.continuesBeTypeTianMen == rr.continuesBeTypeXiaMen && rr.continuesShangMen == rr.continuesBeTypeXiaMen {
				return &netproto.WRTTZ_BetArr{BetArr: betArr}
			}
		}
	}

	return &netproto.WRTTZ_BetArr{BetArr: betArr}
}

func compareVal(a int, b int) bool {
	return a == b
}

func findMax(a int, b int, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		//b > a
		if b > c {
			return b
		} else {
			return c
		}
	}
}

//決定下注幾門
func decideBetNums(a int, b int, c int) int {
	if a == b && a == c {
		return 3
	}
	if (a == b && a > c) || (a == c && a > c) || (b == c && c > c) {
		return 2
	}
	return 1
}

func (rr *DiagonalRobot) ParseResultList(arr []*netproto.WRTTZ_GameResult) {
	arrShangMen := arr[0].GetResult()
	continuesShangMen := checkContinue(arrShangMen)

	if arrShangMen[len(arrShangMen)-1].GetJiang() == BetLose {
		continuesShangMen = 0
	}
	rr.continuesShangMen = continuesShangMen

	arrTianMen := arr[1].GetResult()
	continuesTianMen := checkContinue(arrTianMen)

	if arrTianMen[len(arrTianMen)-1].GetJiang() == BetLose {
		continuesTianMen = 0
	}
	rr.continuesBeTypeTianMen = continuesTianMen

	arrXiaMen := arr[2].GetResult()
	continuesXiaMen := checkContinue(arrXiaMen)

	if arrXiaMen[len(arrXiaMen)-1].GetJiang() == BetLose {
		continuesXiaMen = 0
	}
	rr.continuesBeTypeXiaMen = continuesXiaMen
}

func checkContinue(arr []*netproto.WRTTZ_BetTypeJiang) int {
	maxContinue := 1
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			break
		}
		tmp := arr[i].GetJiang()
		tmp2 := arr[i+1].GetJiang()

		if tmp == netproto.WRTTZ_JiangType_JiangWin && tmp2 == netproto.WRTTZ_JiangType_JiangWin {
			maxContinue++
		} else {
			maxContinue = 1
		}
	}
	return maxContinue
}

func (rr *DiagonalRobot) changeBetCorner(betCorner int32) int32 {
	switch betCorner {
	case BeTypeShangMen:
		return BeTypeXiaJiao
	case BeTypeTianMen:
		return BeTypeGiao
	case BeTypeXiaMen:
		return BeTypeShangJiao
	default:
		return BeTypeShangJiao
	}
}

type ReverseRobot struct {
	roomIndex    int32
	betAmount    int64
	lastJiangArr []int
}

func NewReverseRobot(roomIndex int32) *ReverseRobot {
	return &ReverseRobot{lastJiangArr: make([]int, 3), roomIndex: roomIndex}
}

func (rr *ReverseRobot) ParseParam(params string) {
	betAmount, err := strconv.Atoi(params)
	if err != nil {
		betRange := BetRange[rr.roomIndex]
		rr.betAmount = betRange[len(betRange)]
	} else {
		rr.betAmount = int64(betAmount)
	}
}

func (rr *ReverseRobot) GetBetArr() *netproto.WRTTZ_BetArr {
	betArr := make([]*netproto.WRTTZ_Bet, 0)
	count := 0
	lostIndex := make([]int, 0)

	for i := 0; i < len(rr.lastJiangArr); i++ {
		if rr.lastJiangArr[i] == BetLose {
			count++
			lostIndex = append(lostIndex, i+1)
		}
	}
	switch count {
	case 1:
		betArr = append(betArr, &netproto.WRTTZ_Bet{
			BetType: Int32ToBetType(int32(lostIndex[0])),
			Score:   proto.Int64(rr.betAmount),
		})
		return &netproto.WRTTZ_BetArr{BetArr: betArr}
	case 2:
		betType := getJiao(lostIndex[0], lostIndex[1])
		betArr = append(betArr, &netproto.WRTTZ_Bet{
			BetType: Int32ToBetType(int32(betType)),
			Score:   proto.Int64(rr.betAmount),
		})
		return &netproto.WRTTZ_BetArr{BetArr: betArr}
	case 3,0:
		betType := WeightedRandomS3(ReverseProb)

		switch betType {
		case 0:
			betArr = append(betArr, &netproto.WRTTZ_Bet{
				BetType: Int32ToBetType(int32(BeTypeShangMen)),
				Score:   proto.Int64(rr.betAmount),
			})
		case 1:
			betArr = append(betArr, &netproto.WRTTZ_Bet{
				BetType: Int32ToBetType(int32(BeTypeTianMen)),
				Score:   proto.Int64(rr.betAmount),
			})
		case 2:
			betArr = append(betArr, &netproto.WRTTZ_Bet{
				BetType: Int32ToBetType(int32(BeTypeXiaMen)),
				Score:   proto.Int64(rr.betAmount),
			})
		case 3:

		}


		return &netproto.WRTTZ_BetArr{BetArr: betArr}
	}
	return &netproto.WRTTZ_BetArr{BetArr: betArr}
}

func (rr *ReverseRobot) ParseJiang(arr []*netproto.WRTTZ_BetTypeJiang) {
	for i := 0; i < len(arr); i++ {
		if arr[i].GetBetType() == netproto.WRTTZ_BetType_BeTypeShangMen {
			rr.lastJiangArr[0] = int(JiangType2Int32(arr[i].GetJiang()))
		}
		if arr[i].GetBetType() == netproto.WRTTZ_BetType_BeTypeTianMen {
			rr.lastJiangArr[1] = int(JiangType2Int32(arr[i].GetJiang()))
		}
		if arr[i].GetBetType() == netproto.WRTTZ_BetType_BeTypeXiaMen {
			rr.lastJiangArr[2] = int(JiangType2Int32(arr[i].GetJiang()))
		}
	}
}

func (rr *ReverseRobot) ParseResultList(aee []*netproto.WRTTZ_GameResult) {
}

func getJiao(a int, b int) int {
	switch {
	case a == BeTypeShangMen && b == BeTypeTianMen:
		return BeTypeShangJiao
	case a == BeTypeShangMen && b == BeTypeXiaMen:
		return BeTypeGiao
	case a == BeTypeTianMen && b == BeTypeXiaMen:
		return BeTypeXiaJiao
	}
	return 0
}

type DirectRobot struct {
	roomIndex    int32
	betAmount    int64
	lastJiangArr []int
}

func NewDirectRobot(roomIndex int32) *DirectRobot {
	return &DirectRobot{lastJiangArr: make([]int, 3), roomIndex: roomIndex}
}

func (rr *DirectRobot) ParseParam(params string) {
	betAmount, err := strconv.Atoi(params)
	if err != nil {
		betRange := BetRange[rr.roomIndex]
		rr.betAmount = betRange[len(betRange)]
	} else {
		rr.betAmount = int64(betAmount)
	}
}

func (rr *DirectRobot) GetBetArr() *netproto.WRTTZ_BetArr {
	//金額隨機
	betArr := make([]*netproto.WRTTZ_Bet, 0)
	count := 0
	lostIndex := make([]int, 0)

	for i := 0; i < len(rr.lastJiangArr); i++ {
		if rr.lastJiangArr[i] == BetWin {
			count++
			lostIndex = append(lostIndex, i+1)
		}
	}
	switch count {
	case 1:
		{
			betArr = append(betArr, &netproto.WRTTZ_Bet{
				BetType: Int32ToBetType(int32(lostIndex[0])),
				Score:   proto.Int64(rr.betAmount),
			})
			return &netproto.WRTTZ_BetArr{BetArr: betArr}
		}
	case 2:
		{
			betType := getJiao(lostIndex[0], lostIndex[1])
			betArr = append(betArr, &netproto.WRTTZ_Bet{
				BetType: Int32ToBetType(int32(betType)),
				Score:   proto.Int64(rr.betAmount),
			})
			return &netproto.WRTTZ_BetArr{BetArr: betArr}
		}
	case 3,0:
		{
			betType := WeightedRandomS3(ReverseProb)

			switch betType {
			case 0:
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(int32(BeTypeShangMen)),
					Score:   proto.Int64(rr.betAmount),
				})
			case 1:
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(int32(BeTypeTianMen)),
					Score:   proto.Int64(rr.betAmount),
				})
			case 2:
				betArr = append(betArr, &netproto.WRTTZ_Bet{
					BetType: Int32ToBetType(int32(BeTypeXiaMen)),
					Score:   proto.Int64(rr.betAmount),
				})
			case 3:

			}

			return &netproto.WRTTZ_BetArr{BetArr: betArr}
		}
	}
	return &netproto.WRTTZ_BetArr{BetArr: betArr}
}

func (rr *DirectRobot) ParseJiang(arr []*netproto.WRTTZ_BetTypeJiang) {
	for i := 0; i < len(arr); i++ {
		if arr[i].GetBetType() == netproto.WRTTZ_BetType_BeTypeShangMen {
			rr.lastJiangArr[0] = int(JiangType2Int32(arr[i].GetJiang()))
		}
		if arr[i].GetBetType() == netproto.WRTTZ_BetType_BeTypeTianMen {
			rr.lastJiangArr[1] = int(JiangType2Int32(arr[i].GetJiang()))
		}
		if arr[i].GetBetType() == netproto.WRTTZ_BetType_BeTypeXiaMen {
			rr.lastJiangArr[2] = int(JiangType2Int32(arr[i].GetJiang()))
		}
	}
}

func (rr *DirectRobot) ParseResultList(aee []*netproto.WRTTZ_GameResult) {
}
