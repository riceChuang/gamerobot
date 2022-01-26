package model

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

type GameInfo struct {
	GameID       int32
	Name         string
	RoomType     []string
	Buttons      map[string]*Message
	RoomStrategy map[string]map[string]interface{}
}

type GameMessageInfo struct {
	Handler
	IsButton bool
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
