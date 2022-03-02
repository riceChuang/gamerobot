package wrttz

import (
	"github.com/riceChuang/gamerobot/using/netproto"
	"math/rand"
	"reflect"
)

const (
	MaxBet               = 5
	MustBetModulationVal = 10000
	ParamSplitter        = "|"
	Comma                = ","
	DefaultDiff  = 2
	DefaultTotalAmount  = 10000
)

var BetTypeArray = []int32{BeTypeShangMen, BeTypeTianMen, BeTypeXiaMen, BeTypeShangJiao, BeTypeXiaJiao, BeTypeGiao} //投注區間 莊+上下天門+上下腳+橋

var BetTypeCorner = []int32{BeTypeShangJiao, BeTypeXiaJiao, BeTypeGiao} //上下角橋

var ReverseProb = []float64{1666,1666,1666,5000}

var BetRange = map[int32][]int64{
	1: []int64{200, 500, 1000, 2000, 5000, 10000},               //體驗
	2: []int64{2000, 5000, 10000, 20000, 50000, 100000},         //初級
	3: []int64{20000, 50000, 100000, 200000, 500000, 1000000},   //中級
	4: []int64{50000, 100000, 200000, 500000, 1000000, 2500000}, //高級
}

//map[roomindex]map[nums]int64
var BetStg = map[int32]map[int][][]int64{
	1: {
		1:[][]int64{[]int64{10000}},
		2:[][]int64{
			[]int64{6000,4000},
			[]int64{5000,5000},
			[]int64{8000,2000},
			[]int64{9000,1000},
		},
		3:[][]int64{
			[]int64{5000,4000,1000},
			[]int64{4000,4000,2000},
		},
	},
	2: {
		1:[][]int64{[]int64{100000}},
		2:[][]int64{
			[]int64{60000,40000},
			[]int64{50000,50000},
			[]int64{80000,20000},
			[]int64{90000,10000},
		},
		3:[][]int64{
			[]int64{50000,40000,10000},
			[]int64{40000,40000,20000},
		},
	},
	3: {
		1:[][]int64{[]int64{1000000}},
		2:[][]int64{
			[]int64{600000,400000},
			[]int64{500000,500000},
			[]int64{800000,200000},
			[]int64{900000,100000},
		},
		3:[][]int64{
			[]int64{500000,400000,100000},
			[]int64{400000,400000,200000},
		},
	},
	4: {
		1:[][]int64{[]int64{10000000}},
		2:[][]int64{
			[]int64{6000000,4000000},
			[]int64{5000000,5000000},
			[]int64{8000000,2000000},
			[]int64{9000000,1000000},
		},
		3:[][]int64{
			[]int64{5000000,4000000,1000000},
			[]int64{4000000,4000000,2000000},
		},
	},
}


//下注區間
const (
	BetTypeZhang    = iota //莊
	BeTypeShangMen         //上门
	BeTypeTianMen          //天门
	BeTypeXiaMen           //下门
	BeTypeShangJiao        //上角
	BeTypeXiaJiao          //下角
	BeTypeGiao             //桥
)

func Int32ToBetType(id int32) *netproto.WRTTZ_BetType {
	switch int(id) {
	case BetTypeZhang:
		return netproto.WRTTZ_BetType_BetTypeZhang.Enum()
	case BeTypeShangMen:
		return netproto.WRTTZ_BetType_BeTypeShangMen.Enum()
	case BeTypeTianMen:
		return netproto.WRTTZ_BetType_BeTypeTianMen.Enum()
	case BeTypeXiaMen:
		return netproto.WRTTZ_BetType_BeTypeXiaMen.Enum()
	case BeTypeShangJiao:
		return netproto.WRTTZ_BetType_BeTypeShangJiao.Enum()
	case BeTypeXiaJiao:
		return netproto.WRTTZ_BetType_BeTypeXiaJiao.Enum()
	case BeTypeGiao:
		return netproto.WRTTZ_BetType_BeTypeGiao.Enum()
	}
	return nil
}

type RtpInfo struct {
	GameNo     string
	SysRtp     string
	CurrentRtp int
}

const (
	RobotRand  = iota + 1 //隨機下法
	RobotAllin            //六門全押

	RobotSingleCorner //單角下法
	RobotMultiCorner  //多角下法

	RobotDiagonal //對角下法
	RobotReverse  //逆追路
	RobotDirect   //追路
)



func Corner(id int32) *netproto.WRTTZ_BetType {
	switch int(id) {
	case BetTypeZhang:
		return netproto.WRTTZ_BetType_BetTypeZhang.Enum()
	case BeTypeShangMen:
		return netproto.WRTTZ_BetType_BeTypeShangMen.Enum()
	case BeTypeTianMen:
		return netproto.WRTTZ_BetType_BeTypeTianMen.Enum()
	case BeTypeXiaMen:
		return netproto.WRTTZ_BetType_BeTypeXiaMen.Enum()
	case BeTypeShangJiao:
		return netproto.WRTTZ_BetType_BeTypeShangJiao.Enum()
	case BeTypeXiaJiao:
		return netproto.WRTTZ_BetType_BeTypeXiaJiao.Enum()
	case BeTypeGiao:
		return netproto.WRTTZ_BetType_BeTypeGiao.Enum()
	}
	return nil
}

const (
	BetWin  = 1  //赢
	BetLose = -1 //输
	BetTie  = 0  //和
)


func Int2JiangType(res int32) *netproto.WRTTZ_JiangType {
	switch res {
	case BetWin:
		{
			return netproto.WRTTZ_JiangType_JiangWin.Enum()
		}
	case BetLose:
		{
			return netproto.WRTTZ_JiangType_JiangLose.Enum()
		}
	case BetTie:
		{
			return netproto.WRTTZ_JiangType_JiangTie.Enum()
		}
	}
	return netproto.WRTTZ_JiangType_JiangTie.Enum()
}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}

func JiangType2Int32(jiang netproto.WRTTZ_JiangType) int32 {
	switch jiang {
	case netproto.WRTTZ_JiangType_JiangWin:
		{
			return int32(1)
		}
	case netproto.WRTTZ_JiangType_JiangTie:
		{
			return int32(0)
		}
	case netproto.WRTTZ_JiangType_JiangLose:
		{
			return int32(-1)
		}
	default:
		return int32(99999)
	}
}


func WeightedRandomS3(weights []float64) int {
	n := len(weights)
	if n == 0 {
		return 0
	}
	cdf := make([]float64, n)
	var sum float64 = 0.0
	for i, w := range weights {
		if i > 0 {
			cdf[i] = cdf[i-1] + w
		} else {
			cdf[i] = w
		}
		sum += w
	}
	r := rand.Float64() * sum
	var l, h int = 0, n - 1
	for l <= h {
		m := l + (h-l)/2
		if r <= cdf[m] {
			if m == 0 || (m > 0 && r > cdf[m-1]) {
				return m
			}
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}