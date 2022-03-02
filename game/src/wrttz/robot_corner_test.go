package wrttz

import (
	"github.com/riceChuang/gamerobot/using/netproto"
	"testing"
)

func TestSingleCornerRobot_GetBetArr(t *testing.T) {
	tests := []struct {
		name       string
		Param      string
		roomIndex  int32
		winAmount []int
		lostAmount []int
		currentCorner []int32
		diff    int
		gameResult []*netproto.WRTTZ_BetTypeJiang
		expectCurrentCorner  []int32
		expectWinAmount []int
		expectLostAmount []int
	}{
		{
			name:      "change corner",
			Param:     "2",
			roomIndex: 1,
			winAmount:[]int{2,2,2},
			lostAmount: []int{1,1,1},
			diff:2,
			currentCorner:[]int32{BeTypeShangJiao,BeTypeXiaJiao,BeTypeGiao},
			gameResult: []*netproto.WRTTZ_BetTypeJiang{
				{BetType: Int32ToBetType(BeTypeShangJiao), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaJiao), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeGiao), Jiang: Int2JiangType(BetWin)},
			},
			expectCurrentCorner  :[]int32{
				BeTypeXiaJiao,
				BeTypeGiao,
				BeTypeShangJiao,
			},
			expectWinAmount:[]int{3,3,3},
			expectLostAmount: []int{1,1,1},
		},
		{
			name:      "corner no change",
			Param:     "2",
			roomIndex: 1,
			winAmount:[]int{1,1,1},
			lostAmount: []int{1,1,1},
			diff:10,
			currentCorner:[]int32{BeTypeShangJiao,BeTypeXiaJiao,BeTypeGiao},
			gameResult: []*netproto.WRTTZ_BetTypeJiang{
				{BetType: Int32ToBetType(BeTypeShangJiao), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaJiao), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeGiao), Jiang: Int2JiangType(BetLose)},
			},
			expectCurrentCorner  :[]int32{
				BeTypeXiaJiao,
				BeTypeGiao,
				BeTypeShangJiao,
			},
			expectWinAmount:[]int{2,1,1},
			expectLostAmount: []int{1,2,2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := NewSingleCornerRobot(tt.roomIndex)
			bot.winTime = tt.winAmount
			bot.lostTime = tt.lostAmount
			bot.currentCorners = tt.currentCorner
			bot.diff = tt.diff
			bot.ParseJiang(tt.gameResult)

			for i := 0;i<len(tt.expectLostAmount);i++{
				if tt.expectLostAmount[i] != bot.lostTime[i] {
					t.Errorf("expect lostTime = %v got lostTime = %v",tt.expectLostAmount[i],bot.lostTime[i])
				}
				if tt.expectWinAmount[i] != bot.winTime[i] {
					t.Errorf("expect winTime  = %v got winTime  = %v",tt.expectWinAmount[i],bot.winTime[i])
				}
			}
			bot.GetBetArr()
			for i := 0;i<len(tt.expectCurrentCorner);i++{
				if tt.currentCorner[i] != tt.expectCurrentCorner[i]{
					t.Errorf("expect currentCorner  = %v got currentCorner  = %v",tt.expectCurrentCorner[i],tt.currentCorner[i] )
				}
			}
		})
	}
}


func TestMultiCornerRobot_GetBetArr(t *testing.T) {
	tests := []struct {
		name       string
		Param      string
		roomIndex  int32
	}{
		{
			name:"test1",
			Param: "10000",
			roomIndex: 2,
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			bot := NewMultiCornerRobot(tt.roomIndex)
			bot.ParseParam(tt.Param)
			bot.GetBetArr()
		})
	}
}