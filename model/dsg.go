package model

type DSGLoginReq struct {
	LoginDomain string
	AgentID     int
	Account     string
	Password    string
	GameID      int32
}

type DSGLoginResp struct {
	URL      string
	UserName string
	Token    string
}

type DSGStoreMoneyReq struct {
	LoginDomain string
	AgentID     int
	Account     string
	Money       int32
}


type DSGStoreMoneyResp struct {
	Money int32
}