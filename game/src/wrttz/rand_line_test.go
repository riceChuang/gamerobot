package wrttz

import (
	"github.com/riceChuang/gamerobot/using/netproto"
	"testing"
)

func Test_checkContinue(t *testing.T) {
	tests := []struct {
		name string
		arr  []*netproto.WRTTZ_BetTypeJiang
		wantVal  int
	}{
		{
			name: "test1",
			arr: []*netproto.WRTTZ_BetTypeJiang{
				{BetType: Int32ToBetType(BeTypeShangMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeTianMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
			},
			wantVal:1,
		},
		{
			name: "test2",
			arr: []*netproto.WRTTZ_BetTypeJiang{
				{BetType: Int32ToBetType(BeTypeTianMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeShangMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeTianMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
			},
			wantVal:2,
		},
		{
			name: "test2",
			arr: []*netproto.WRTTZ_BetTypeJiang{
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeTianMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeShangMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeTianMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
			},
			wantVal:2,
		},
		{
			name: "test2",
			arr: []*netproto.WRTTZ_BetTypeJiang{
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeTianMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeShangMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeTianMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
			},
			wantVal:1,
		},
		{
			name: "test2",
			arr: []*netproto.WRTTZ_BetTypeJiang{
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetLose)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
				{BetType: Int32ToBetType(BeTypeXiaMen), Jiang: Int2JiangType(BetWin)},
			},
			wantVal:4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkContinue(tt.arr)
			if got != tt.wantVal{
				t.Errorf("want val = %v got val = %v",tt.wantVal,got)
			}
		})
	}
}

func Test_decideBetNums(t *testing.T){
	tests := []struct {
		name string
		a int
		b int
		c int
		wantVal  int
	}{
		{
			name:"test1",
			a:1,
			b:1,
			c:1,
			wantVal:3,
		},
		{
			name:"test1",
			a:2,
			b:1,
			c:1,
			wantVal:1,
		},
		{
			name:"test1",
			a:2,
			b:2,
			c:0,
			wantVal:2,
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			got:= decideBetNums(tt.a,tt.b,tt.c)
			if got != tt.wantVal{
				t.Errorf("got val %v,want val %v",got,tt.wantVal)
			}
		})
	}
}