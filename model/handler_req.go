package model

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	GameID   int32  `json:"gameid"`
	RoomType string `json:"roomtype"`
	Env      string `json:"env"`
}

type StressReq struct {
	GameID     int32  `json:"gameid"`
	RoomType   string `json:"roomtype"`
	Env        string `json:"env"`
	RobotCount int    `json:"robotcount"`
	Strategy   string `json:"strategy"`
}

type StrategyItem struct {
	BetOdd   string `json:"betodd"`
	Params     string `json:"params"`
}
