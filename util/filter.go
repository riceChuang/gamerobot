package util

import (
	"github.com/riceChuang/gamerobot/model"
)

type SendInfo struct {
	GameSendMap map[int32]*SendFilter
}

type SendFilter struct {
	SendType      int
	sendAllClient bool           //如果沒寫就全送
	sendAllGame   bool           //如果沒寫就全送
	gameSidMap    map[int32]bool //game robot要送什麼id
	clientSidMap  map[int32]bool //client robot要送什麼id
}

func NewSendFilter(cfg []model.CommonCfg) *SendInfo {
	sendInfo := make(map[int32]*SendFilter)
	for _, game := range cfg {
		sendFilter := &SendFilter{}
		gameMap := make(map[int32]bool)
		for _, gameSids := range game.GameSIds {
			gameMap[gameSids] = true
		}
		sendFilter.gameSidMap = gameMap

		clientMap := make(map[int32]bool)
		for _, clientSids := range game.ClientSIds {
			clientMap[clientSids] = true
		}
		sendFilter.clientSidMap = clientMap

		if len(gameMap) == 0 {
			sendFilter.sendAllGame = true
		}
		if len(clientMap) == 0 {
			sendFilter.sendAllClient = true
		}
		sendInfo[game.GameID] = sendFilter
	}
	return &SendInfo{GameSendMap: sendInfo}
}

//raw msg 轉乘client 的msg
func (s *SendFilter) FilterClientMsg(msg *model.WSMessage) bool {
	sId := msg.Msg.SClassID
	_, exists := s.clientSidMap[sId]

	if s.sendAllClient == true || exists {
		return true
	}
	return false
}

func (s *SendFilter) FilterGameMsg(msg *model.WSMessage) bool {
	sId := msg.Msg.SClassID
	_, exists := s.gameSidMap[sId]
	if exists || s.sendAllGame {
		return true
	} else {
		return false
	}
}
