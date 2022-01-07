package model

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	GameID   int32  `json:"gameid"`
	RoomType string `json:"roomtype"`
	Env      string `json:"env"`
}

