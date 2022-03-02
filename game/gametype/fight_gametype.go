package gametype

import (
	"github.com/riceChuang/gamerobot/common"
	connection "github.com/riceChuang/gamerobot/framework/connection/connect"
	"github.com/riceChuang/gamerobot/framework/connection/connecttype"
	"github.com/riceChuang/gamerobot/framework/gamehandler"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/usecase"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util/config"
)

// Base of stragedy
type FightBase struct {
	*BaseGameType
}

// NewBase ...
func NewFightBase(cfg model.CommonCfg, gamelogic gamehandler.GameLogicBase) GameType {
	fb := &FightBase{
		BaseGameType: NewBaseGameType(cfg, gamelogic),
	}

	return fb
}

func (fb *FightBase) Initialize(gameConnect *connecttype.GameConnect) {
	fb.logger.Info("初始化 fightBase")
	if gameConnect != nil {
		fb.gameConnect = gameConnect
		gameConnect.GameWS.Register(&model.Handler{
			BClassID:  int32(netproto.MessageBClassID_Game),
			SClassID:  int32(netproto.GameMessageClassID_UserStateChangeID),
			MsgType:   &netproto.UserStateChange{},
			OnMessage: fb.onUserStateChange,
		})
	}
}

func (fb *FightBase) HandleMessage(message *model.WSMessage) {

	switch message.Msg.SClassID {
	default:
		fb.gameHandler.HandleMessage(message)
	}
	return
}

func (fb *FightBase) onUserStateChange(msg interface{}) {
	data, ok := msg.(*netproto.UserStateChange)
	if !ok {
		fb.logger.Error("onUserStateChange Error: Data Type Wrong!")
		return
	}

	user := data.GetUser()
	if user.GetUserID() == fb.gameConnect.GetUserID() {
		fb.gameConnect.UserMoney = int32(user.GetScore() / 100)
		fb.logger.Info("onUserStateChange User update money: %d", fb.gameConnect.UserMoney)
		fb.logger.Info("onUserStateChange User current stateValue: %d", user.GetStateValue())
		cfg := config.GetGameInstanceByID(fb.GetGameHandler().GetGameRoomIndex())
		//少於最小金額
		if fb.gameConnect.UserMoney < cfg.MinMoney {
			fb.gameConnect.CleanGs()
			usecase := usecase.NewUseCase()
			dsgStoreReq := &model.DSGStoreMoneyReq{
				LoginDomain: fb.gameConnect.Env.LoginDomain,
				AgentID:     fb.gameConnect.Env.AgentID,
				Account:     user.GetNickName(),
				Money:       cfg.StoredMoney,
			}
			fb.logger.Errorf("[沒錢了沒錢了沒錢了] 玩家:%v, 沒錢了, 當前金額:%v, 需要金額:%v, 儲值金額:%v", user.GetNickName(), fb.gameConnect.UserMoney, cfg.MinMoney, cfg.StoredMoney)
			dsgStoreResp, err := usecase.DsgAPIBase.StoreMoney(dsgStoreReq)
			if err != nil {
				fb.logger.Error(err)
				return
			}
			fb.logger.Infof("沒錢了沒錢了沒錢了 money:%v", dsgStoreResp.Money)
			fb.gameConnect.UserMoney = dsgStoreResp.Money
			nProtoConn := connection.NewProtoConnect(connection.NewConn(fb.gameConnect.GameWS.GetURL(), nil, common.GameConnect))
			fb.gameConnect.GameWS = nProtoConn
			fb.Initialize(fb.gameConnect)
			fb.gameConnect.ConnectGameWs()
			return
		}

		if user.GetStateValue() == int32(common.StateValueSit) {
			fb.sendReady()
		}
	}
}

func (fb *FightBase) sendReady() {
	if fb.gameConnect != nil {
		fb.logger.Info("[我準備了！！！！！！！！]")
		fb.gameConnect.Request(
			fb.gameConnect.GameWS,
			int32(netproto.MessageBClassID_GameRoom),
			int32(netproto.GameRoomClassID_GameReadyID),
			nil,
		)
	}
	return
}
