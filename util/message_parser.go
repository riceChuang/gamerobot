package util

import (
	"errors"
	"github.com/riceChuang/gamerobot/model"
)

// Parser ...
type Parser interface {
	UnMarshal(bytes []byte) (*model.Message, error)
	Marshal(msg *model.Message) ([]byte, error)
}

//ByteParser 基于[]byte消息内容的消息解析器
type ByteParser struct {
}

// Marshal convert Message to bytes
func (bmp *ByteParser) Marshal(msg *model.Message) ([]byte, error) {
	b1 := IntToBytes(int32(msg.BClassID), true)
	b2 := IntToBytes(int32(msg.SClassID), true)

	b3, ok := msg.Data.([]byte)

	if !ok {
		return nil, errors.New("数据类型必须为[]byte")
	}

	return append(append(b1, b2...), b3...), nil
}

// UnMarshal convert bytes to Message
func (bmp *ByteParser) UnMarshal(data []byte) (*model.Message, error) {
	msg := new(model.Message)
	length := BytesToInt(data[:4], true)
	msg.BClassID = BytesToInt(data[4:8], true)
	msg.SClassID = BytesToInt(data[8:12], true)

	msg.Data = data[12 : length+4]

	return msg, nil
}
