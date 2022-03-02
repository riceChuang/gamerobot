package common

type GameServerID int32

const (
	GameID_QZNN   GameServerID = 2  //搶莊牛牛
	GameID_ZJH    GameServerID = 3  //喜錢炸金花
	GameID_TBNN   GameServerID = 7  //通比牛牛
	GameID_SG     GameServerID = 8  //三公
	GameID_SSS    GameServerID = 10 //十三水
	GameID_LHD    GameServerID = 11 //龍虎鬥
	GameID_BJL    GameServerID = 13 //百家樂
	GameID_BRNN   GameServerID = 19 //百人牛牛
	GameID_HHDZ   GameServerID = 20 //红黑大战
	GameID_CD21   GameServerID = 22 //21點
	GameID_WRJH   GameServerID = 28 //萬人炸金花
	GameID_BAR28  GameServerID = 29 //28槓
	GameID_QZPN   GameServerID = 30 //搶莊牌九
	GameID_CPQN   GameServerID = 31 //看三搶裝牛牛
	GameID_CPQN4  GameServerID = 32 //看四搶莊牛牛
	GameID_IDAZJH GameServerID = 33 //印度炸金花
	GameID_BCBM   GameServerID = 34 //奔馳寶馬
	GameID_FQZS   GameServerID = 35 //飞禽走兽
	GameID_FP     GameServerID = 36 //森林舞會
	GameID_BYDH   GameServerID = 38 //捕魚大亨
	GameID_BRSB   GameServerID = 40 //百人骰宝
	GameID_WRTTZ  GameServerID = 41 //万人推筒子
)

var (
	GameServerID_EngName = map[GameServerID]string{
		GameID_QZNN:   "QZNN",
		GameID_ZJH:    "SCZJH",
		GameID_TBNN:   "TBNN",
		GameID_SG:     "SG",
		GameID_SSS:    "SSS",
		GameID_LHD:    "LHD",
		GameID_BJL:    "BJL",
		GameID_BRNN:   "BRNN",
		GameID_HHDZ:   "HHDZ",
		GameID_CD21:   "CD21",
		GameID_WRJH:   "WRJH",
		GameID_BAR28:  "BAR28",
		GameID_QZPN:   "QZPN",
		GameID_CPQN:   "CPQN",
		GameID_CPQN4:  "CPQN4",
		GameID_IDAZJH: "IDAZJH",
		GameID_BCBM:   "BCBM",
		GameID_FQZS:   "FQZS",
		GameID_FP:     "FP",
		GameID_BYDH:   "BYDH",
		GameID_BRSB:    "BRSB",
		GameID_WRTTZ:  "WRTTZ",
	}

	// GameIDtoName ...
	GameServerID_Name = map[GameServerID]string{
		GameID_QZNN:   "搶裝牛牛",
		GameID_ZJH:    "喜錢炸金花",
		GameID_TBNN:   "通比牛牛",
		GameID_SG:     "三公",
		GameID_SSS:    "十三水",
		GameID_LHD:    "龍虎鬥",
		GameID_BJL:    "百家樂",
		GameID_BRNN:   "百人牛牛",
		GameID_HHDZ:   "红黑大战",
		GameID_CD21:   "21點",
		GameID_WRJH:   "萬人炸金花",
		GameID_BAR28:  "28槓",
		GameID_QZPN:   "搶莊牌九",
		GameID_CPQN:   "看三搶裝牛牛",
		GameID_CPQN4:  "看四搶莊牛牛",
		GameID_IDAZJH: "印度炸金花",
		GameID_BCBM:   "奔馳寶馬",
		GameID_FQZS:   "飞禽走兽",
		GameID_FP:     "森林舞會",
		GameID_BYDH:   "捕魚大亨",
		GameID_BRSB:    "百人骰宝",
		GameID_WRTTZ:  "万人推筒子",
	}
	// GameNameToID ...
	GameNameToID = map[string]int{
		"QZNN":    2,
		"TBNN":    7,
		"SG":      8,
		"SSS":     10,
		"LHD":     11,
		"BJL":     13,
		"ERMJ":    17,
		"BRNN":    19,
		"HHDZ":    20,
		"CD21":    22,
		"WRJH":    28,
		"ERG":     29,
		"PG":      30,
		"KPQZNN":  31,
		"K4PQZNN": 32,
		"FP":      36,
		"RU":      37,
		"BYDH":    38,
		"DTNN":    39,
		"BRSB":    40,
		"WRTTZ":   41,
	}
	GameIdToInt32 = map[GameServerID]int32{
		GameID_QZNN:   2,  //搶莊牛牛
		GameID_ZJH:    3,  //喜錢炸金花
		GameID_TBNN:   7,  //通比牛牛
		GameID_SG:     8,  //三公
		GameID_SSS:    10, //十三水
		GameID_LHD:    11, //龍虎鬥
		GameID_BJL:    13, //百家樂
		GameID_BRNN:   19, //百人牛牛
		GameID_HHDZ:   20, //红黑大战
		GameID_CD21:   22, //21點
		GameID_WRJH:   28, //萬人炸金花
		GameID_BAR28:  29, //28槓
		GameID_QZPN:   30, //搶莊牌九
		GameID_CPQN:   31, //看三搶裝牛牛
		GameID_CPQN4:  32, //看四搶莊牛牛
		GameID_IDAZJH: 33, //印度炸金花
		GameID_BCBM:   34, //奔馳寶馬
		GameID_FQZS:   35, //飞禽走兽
		GameID_FP:     36, //森林舞會
		GameID_BYDH:   38, //捕魚大亨
		GameID_BRSB:   40, //捕魚大亨
		GameID_WRTTZ:  41, //万人推筒子
	}
)

func GameServerIDToString(id GameServerID) string {
	return GameServerID_Name[id]
}

// Env ...
type Env string

const (
	DEV     Env = "DEV"
	QA      Env = "QA"
	STAGE   Env = "STAGE"
	RELEASE Env = "RELEASE"
	FEATURE Env = "FEATURE"
	RTP     Env = "RTP"
)

type StateValue int32

const (
	// StateValueFree 用户站立状态
	StateValueFree StateValue = 0
	// StateValueQueue 用户排队状态
	StateValueQueue StateValue = 1
	// StateValueSit 用户坐下状态
	StateValueSit StateValue = 2
	// StateValueReady 用户举手状态
	StateValueReady StateValue = 3
	// StateValueGame 用户游戏状态
	StateValueGame StateValue = 4
	// StateValueOffline 用户断线状态
	StateValueOffline StateValue = 5
)
