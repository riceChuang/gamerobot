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
)

var (
	GameServerID_Name = map[GameServerID]string{
		GameID_QZNN:   "搶莊牛牛",
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
