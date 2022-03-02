package wrttz

import (
	"github.com/riceChuang/gamerobot/using/netproto"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type RandRobot struct {
	weights         []int64
	betRangeByIndex []int64
}

func NewRandRobot(roomIndex int32) *RandRobot {
	rand.Seed(time.Now().UnixNano())
	return &RandRobot{betRangeByIndex: BetRange[roomIndex]}
}

// "1,1000|2,2000|3,5000|4,6000|5,1000|6,2000|7,5000|8,6000"
func (rr *RandRobot) ParseParam(params string) {
	weights := make([]int64, 0)
	paramArr := strings.Split(params, ParamSplitter)
	for i := 0; i < len(paramArr); i++ {
		valStr := strings.Split(paramArr[i], Comma)[1]
		wets, err := strconv.Atoi(valStr)
		if err != nil {
			continue
		}
		weights = append(weights, int64(wets))
	}
	rr.weights = weights
}

func (rr *RandRobot) GetBetArr() *netproto.WRTTZ_BetArr {
	betArr := make([]*netproto.WRTTZ_Bet, 0)
	for i := 0; i < len(rr.weights); i++ {
		doBet := DecideBet(rr.weights[i])
		betType := Int32ToBetType(int32(i + 1))
		if doBet {
			betAmount := rr.betRangeByIndex[rand.Intn(len(BetRange))]
			betArr = append(betArr, &netproto.WRTTZ_Bet{BetType: betType, Score: &betAmount})
		}
	}
	//決定是否下注
	return &netproto.WRTTZ_BetArr{BetArr: betArr}
}

func (rr *RandRobot) ParseJiang([]*netproto.WRTTZ_BetTypeJiang){

}

func (rr *RandRobot)ParseResultList(aee []*netproto.WRTTZ_GameResult){

}

//[ 1000,2000,3000,4000,5000,6000,7000,8000,9000,10000 ]
func DecideBet(weight int64) bool {
	tmp := rand.Intn(10000)
	if tmp >= 1 && int64(tmp) <= int64(weight) {
		return true
	}
	return false
}

type AllInRobot struct {
	BetAmount int64
}

func NewAllInRobot(roomIndex int32) *AllInRobot {
	rand.Seed(time.Now().UnixNano())
	//取最高注
	arr := BetRange[roomIndex]
	return &AllInRobot{BetAmount: arr[len(arr)-1]}
}

func (rr *AllInRobot) ParseParam(params string) {
}

func (rr *AllInRobot) GetBetArr() *netproto.WRTTZ_BetArr {
	betArr := make([]*netproto.WRTTZ_Bet, 0)
	for i := 0; i < len(BetTypeArray); i++ {
		betType := Int32ToBetType(int32(i + 1))
		amount := rr.BetAmount
		betArr = append(betArr, &netproto.WRTTZ_Bet{BetType: betType, Score: &amount})
	}
	//決定是否下注
	return &netproto.WRTTZ_BetArr{BetArr: betArr}
}

func (rr *AllInRobot) ParseJiang([]*netproto.WRTTZ_BetTypeJiang){

}

func (rr *AllInRobot)ParseResultList(aee []*netproto.WRTTZ_GameResult){

}