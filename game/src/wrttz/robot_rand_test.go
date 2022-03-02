package wrttz

import (
	"testing"
)

func TestRandRobot_GetBetArr(t *testing.T) {
	tests := []struct{
		name string
		Param string
		roomIndex int32
	}{
		{
			name:"test1",
			Param: `{"betodd":"1","params":"1,1000|2,2000|3,5000|4,6000|5,1000|6,2000|7,5000|8,6000"}`,
			roomIndex: int32(1),
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			bot := NewRandRobot(tt.roomIndex)
			bot.ParseParam(tt.Param)
			res := bot.GetBetArr()
			t.Logf("bet arr = %v\n",res)
		})
	}
}

