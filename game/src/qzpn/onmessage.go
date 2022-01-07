package qzpn

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util"
)

func (qzpn *QZPNLogic) hiOne(data interface{}) {
	qzpn.Logger.Info("Hi ONE ONE ONE ONE")
}



func (qzpn *QZPNLogic) hiTwo(data interface{}) {
	qzpn.Logger.Info("Hi TWO TWO TWO TWO")
}


func (qzpn *QZPNLogic) hiThree(data interface{}) {
	msg := data.(*model.WSMessage)

	msg.From = common.Game
	msg.To = common.ClientSender
	msg.Msg.Data = "我收到了訊息但我沒傳給遊戲"
	util.GetGameDispatcher().AddMessage(msg)
}