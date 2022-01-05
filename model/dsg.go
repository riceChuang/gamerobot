package model

import "github.com/riceChuang/gamerobot/using/netproto"

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
	UserInfo *netproto.UserLoginRet
}
