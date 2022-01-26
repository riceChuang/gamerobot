package gamehandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/config"
	"github.com/riceChuang/gamerobot/util/logs"
	"sync"
)

type GameLogicBase interface {
	GetMessageBtn() map[string]*model.Message                       //取得串接按鈕
	HandleMessage(message *model.WSMessage)                         //處理訊息
	UnmarshalMessage(message *model.Message) (proto.Message, error) //解析消息
	GetGameRoomIndex() string                                       //取得遊戲id, e.x. qzpn-1
	SetUserMoney(money int32)                                       //同步玩家金額
	SetStrategy(data interface{}) error                             //同步遊戲策略
}

// Base
type Base struct {
	MessageHandler sync.Map
	Logger         *log.Entry
	Transfer       func(int32) string
	GameID         common.GameServerID
	RoomIndex      int32
	GameName       string
	UserMoney      int32
}

// NewBase ...
func NewBase(gameid common.GameServerID, roomIndex int32) *Base {
	b := &Base{
		MessageHandler: sync.Map{},
		Logger: logs.GetLogger().WithFields(log.Fields{
			"server": common.GameServerID_EngName[gameid],
		}),
		GameID:    gameid,
		RoomIndex: roomIndex,
		GameName:  common.GameServerID_EngName[gameid],
	}
	b.Transfer = func(i int32) string {
		return ""
	}
	return b
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
			Butttons[b.Transfer(handler.SClassID)] = &model.Message{
				BClassID: handler.BClassID,
				SClassID: handler.SClassID,
			}
		}
		return true
	})
	return Butttons
}

//get gameindex 30-1
func (b *Base) GetGameRoomIndex() string {
	return fmt.Sprintf("%v-%v", b.GameID, b.RoomIndex)
}

//set empty
func (b *Base) SetStrategy(data interface{}) error {
	return nil
}

//set money
func (b *Base) SetUserMoney(money int32) {
	b.UserMoney = money
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

//unmarshal room config for game struct
func (b *Base) UnmarshalRoomConfig(roomID string, v interface{}) error {
	//id組合方法 GameID-RoomIndex e.x. 30-1
	roomCfg := config.GetGameInstanceByID(roomID)
	if roomCfg == nil {
		return errors.New("cant find roomid")
	}
	return nil
}

func (b *Base) UnmarshalConfig(data interface{}, result interface{}) error {
	// Convert map to json string
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Convert json string to struct
	if err = json.Unmarshal(jsonStr, result); err != nil {
		return err
	}
	return nil
}

func (b *Base) sendMessage(m *model.WSMessage, h *model.GameMessageInfo) {
	b.Logger.Println("dispatch message", m.Msg.GetName())
	if len(m.Msg.Data.([]byte)) == 0 && h.OnMessage != nil {
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
