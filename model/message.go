package model

import (
	"fmt"
	"github.com/riceChuang/gamerobot/using/netproto"
	"google.golang.org/protobuf/proto"
)

var (
	BToS = map[int32]map[int32]string{
		int32(netproto.MessageBClassID_Common):        netproto.PlatformCommonClassID_name,
		int32(netproto.MessageBClassID_DBServer):      netproto.DBServerClassID_name,
		int32(netproto.MessageBClassID_GameRoom):      netproto.GameRoomClassID_name,
		int32(netproto.MessageBClassID_ServerMgm):     netproto.ServerMgmClassID_name,
		int32(netproto.MessageBClassID_Game):          netproto.GameMessageClassID_name,
		int32(netproto.MessageBClassID_Hall):          netproto.HallMsgClassID_name,
		int32(netproto.MessageBClassID_NotifyServer):  netproto.NotifyServerClassID_name,
		int32(netproto.MessageBClassID_NewByDBServer): netproto.DBServerClassID_name,
	}
)

// Message 网络消息对象定义
type Message struct {
	BClassID int32       //一级协议号
	SClassID int32       //二级协议号
	Data     interface{} //消息数据
}

// New a Message
func New(bclsID int32, sclsID int32, data interface{}) *Message {
	msg := new(Message)
	msg.BClassID = bclsID
	msg.SClassID = sclsID
	msg.Data = data

	return msg
}

func (msg *Message) String() string {
	return fmt.Sprintf("%v:%v", msg.BClassID, msg.SClassID)
}

func (msg *Message) GetName() string {
	return fmt.Sprintf("[%v] [%v:%v]  ", msg.String(), netproto.MessageBClassID_name[msg.BClassID], BToS[msg.BClassID][msg.SClassID])
}

// Equals 消息ID是否一致
func (msg *Message) Equals(othermsg *Message) bool {
	return msg.BClassID == othermsg.BClassID && msg.SClassID == othermsg.SClassID
}

// NewProto accept protoData for input
func NewProto(bclsID int32, sclsID int32, protoData interface{}) (*Message, error) {
	var bData []byte
	var err error
	if protoData != nil {
		protoMsg, ok := protoData.(proto.Message)
		if !ok {
			return nil, fmt.Errorf("Not correct protoData when new a proto message")
		}

		bData, err = proto.Marshal(protoMsg)
		if err != nil {
			return nil, err
		}
	}

	message := New(
		bclsID,
		sclsID,
		bData,
	)

	return message, nil
}
