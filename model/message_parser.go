package model

import (
	"errors"
	"github.com/riceChuang/gamerobot/util"
)

// Parser ...
type Parser interface {
	UnMarshal(bytes []byte) (*Message, error)
	Marshal(msg *Message) ([]byte, error)
}

//ByteParser 基于[]byte消息内容的消息解析器
type ByteParser struct {
}

// Marshal convert Message to bytes
func (bmp *ByteParser) Marshal(msg *Message) ([]byte, error) {
	b1 := util.IntToBytes(int32(msg.BClassID), true)
	b2 := util.IntToBytes(int32(msg.SClassID), true)

	b3, ok := msg.Data.([]byte)

	if !ok {
		return nil, errors.New("数据类型必须为[]byte")
	}

	return append(append(b1, b2...), b3...), nil
}

// UnMarshal convert bytes to Message
func (bmp *ByteParser) UnMarshal(data []byte) (*Message, error) {
	msg := new(Message)
	length := util.BytesToInt(data[:4], true)
	msg.BClassID = util.BytesToInt(data[4:8], true)
	msg.SClassID = util.BytesToInt(data[8:12], true)

	msg.Data = data[12 : length+4]

	return msg, nil
}