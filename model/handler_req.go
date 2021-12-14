package model

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	GameID   int32  `json:"gameid"`
	Env      int32  `json:"env"`
}


// Handler handle message
type Handler struct {
	BClassID  int32 //一级协议号
	SClassID  int32 //二级协议号
	MsgType   proto.Message
	OnMessage func(data interface{})
}

func (h *Handler) String() string {
	return fmt.Sprintf("%d:%d", h.BClassID, h.SClassID)
}
