package gamehandler

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sync"
)

type GameLogicBase interface {
	GetMessageBtn() map[string]*model.Message
	HandleMessage(message *model.WSMessage)
	UnmarshalMessage(message *model.Message) (proto.Message, error)
}

// Base
type Base struct {
	MessageHandler sync.Map
	Logger         *log.Entry
	transfer       func(int32) string
}

// NewBase ...
func NewBase(gameName string, tf func(int32) string) *Base {
	return &Base{
		MessageHandler: sync.Map{},
		Logger: logs.GetLogger().WithFields(log.Fields{
			"server": gameName,
		}),
		transfer: tf,
	}
}

func (b *Base) HandleMessage(message *model.WSMessage) {
	if handler, k := b.MessageHandler.Load(message.Msg.String()); k {
		h := handler.(*model.GameMessageInfo)
		b.sendMessage(message, h)
		return
	}
}

func (b *Base) UnmarshalMessage(message *model.Message) (proto.Message, error) {
	if handler, k := b.MessageHandler.Load(message.String()); k {
		h := handler.(*model.GameMessageInfo)
		if h.MsgType != nil {
			if err := proto.Unmarshal(message.Data.([]byte), h.MsgType); err != nil {
				b.Logger.Error("client proto unmarshal error: ", err.Error(), fmt.Sprintf("%+v", message.Data))
				return nil, err
			}
			var resp = h.MsgType
			return resp, nil
		}
		b.Logger.Errorf("not found unmarshal handler:%v", message.String())
		return nil, errors.New(fmt.Sprintf("not found unmarshal MsgType:%v", message.String()))
	}
	return nil, errors.New(fmt.Sprintf("not found unmarshal handler:%v", message.String()))
}

func (b *Base) GetMessageBtn() map[string]*model.Message {
	var Butttons = make(map[string]*model.Message)
	b.MessageHandler.Range(func(key, value interface{}) bool {
		handler := value.(*model.GameMessageInfo)
		if handler.IsButton {
			Butttons[b.transfer(handler.SClassID)] = &model.Message{
				BClassID: handler.BClassID,
				SClassID: handler.SClassID,
			}
		}
		return true
	})
	return Butttons
}

// Register Handler ...
func (b *Base) Register(handlers ...*model.GameMessageInfo) error {
	for _, h := range handlers {
		if _, exist := b.MessageHandler.Load(h.String()); exist {
			return fmt.Errorf("handler already exist: %s", h.String())
		}
		b.Logger.Debug("Register handler", h.String())
		b.MessageHandler.Store(h.String(), h)
	}
	return nil
}

func (b *Base) sendMessage(m *model.WSMessage, h *model.GameMessageInfo) {
	b.Logger.Println("dispatch message", m.Msg.GetName())
	if m.Msg.Data == nil && h.OnMessage != nil {
		//for button message
		h.OnMessage(m)
		return
	} else if h.MsgType != nil {
		if err := proto.Unmarshal(m.Msg.Data.([]byte), h.MsgType); err != nil {
			b.Logger.Println("client proto unmarshal error: ", err.Error(), fmt.Sprintf("%+v", m.Msg.Data))
			return
		}
		m.Msg.Data = h.MsgType
		if h.OnMessage != nil {
			h.OnMessage(m)
		}
		return
	}
	b.Logger.Error("找不到 msgType")
}
