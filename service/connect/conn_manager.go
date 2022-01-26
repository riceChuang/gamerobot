package connect

import (
	"context"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework/connection/connect"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util"
	"github.com/riceChuang/gamerobot/util/logs"
	"net/http"
	"sync"
)

var clientGameCommunicateOnce sync.Once
var clientGameCommunicateManager *ClientGameCommunicateManager

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ClientGameCommunicateManager struct {
	mu           *sync.Mutex
	ctx          context.Context
	cancel       context.CancelFunc
	wsSeverAddr  string
	client       map[string]*ClientConn
	logger       *log.Entry
	parser       util.Parser
	gameWSUrlMap map[string]string
	hallWS       *connect.ProtoConnect
}

func NewClientWsToGameServer(wsAddr string) *ClientGameCommunicateManager {
	clientGameCommunicateOnce.Do(func() {
		clientGameCommunicateManager = &ClientGameCommunicateManager{
			mu:          &sync.Mutex{},
			wsSeverAddr: wsAddr,
			client:      make(map[string]*ClientConn),
			logger: logs.GetLogger().WithFields(log.Fields{
				"server": "connection_manager",
			}),
			parser: &util.ByteParser{},
		}

		util.GetClientDispatcher().AddMsgPasser(string(common.ClientServerTransfer), clientGameCommunicateManager.ClientServerTransferMessage)
		util.GetClientDispatcher().AddMsgPasser(string(common.ClientSender), clientGameCommunicateManager.ClientSender)
		util.GetGameDispatcher().AddMsgPasser(string(common.GameServerTransfer), clientGameCommunicateManager.GameServerTransferMessage)
		util.GetGameDispatcher().AddMsgPasser(string(common.GameSender), clientGameCommunicateManager.GameSender)
	})
	return clientGameCommunicateManager
}

func GetClientGameCommunicateManager() *ClientGameCommunicateManager {
	return clientGameCommunicateManager
}

// ClientServerTransferMessage 轉換從client端來資料,處理後丟給遊戲
func (c2gm *ClientGameCommunicateManager) ClientServerTransferMessage(msg *model.WSMessage) {
	//轉換或處理資料
	/*
		1. 依照看是哪個game的資訊找尋相對應的game_factory 將訊息丟給工廠處理
		2. 工廠若處理完畢會再送往傳送game的通道
	*/
	c2gm.logger.Infof("收到訊息 處理func:TransferClientMessage, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Msg.String())
	//特殊 前端client ws斷線時
	if msg.Msg.BClassID == common.WSClose {
		c2gm.mu.Lock()
		for _, c := range c2gm.client {
			if c.gameConn != nil {
				err := c.CloseProtoConnect()
				if err != nil {
					c2gm.logger.Error(err)
					return
				}
			}
		}
	}

	client := c2gm.GetClient(msg.ClientID)
	if client == nil {
		return
	}

	if !client.IsGameConnAlive() {
		c2gm.logger.Infof("gamews is not alive")
		return
	}

	client.GameType.HandleMessage(msg)
}

// GameServerTransferMessage 轉換game端來的資料,處理後丟給client
func (c2gm *ClientGameCommunicateManager) GameServerTransferMessage(msg *model.WSMessage) {
	//轉換或處理資料
	/*
		1. 需要把拿game的資訊丟一份往client顯示
		2. 依照看是哪個game的資訊找尋相對應的game_factory 將訊息丟給工廠處理
		3. 工廠處理完訊息可能會有部分訊息網game通道送 有可能也會網client送
	*/
	c2gm.logger.Infof("收到訊息 處理func:TransferGameMessage, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Msg.String())
	client := c2gm.GetClient(msg.ClientID)
	if client == nil {
		return
	}
	//也將資料丟給client顯示
	msgData := *msg.Msg
	toClientMsg := &model.WSMessage{
		From:     common.GameServerTransfer,
		To:       common.ClientSender,
		ClientID: msg.ClientID,
		Msg:      &msgData,
	}
	//將資料丟給game 工廠處理
	client.GameType.HandleMessage(msg)

	//轉換資料
	DetailData, err := client.GameType.UnmarshalMessage(&msgData)
	if err != nil {
		c2gm.logger.Error(err)
		return
	}
	toClientMsg.Msg.Data = DetailData
	c2gm.logger.Infof("[送送送送送], from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Msg.String())
	util.GetClientDispatcher().AddMessage(toClientMsg)
}

// GameSender 找到client 的ws 將資料ws寫給他
func (c2gm *ClientGameCommunicateManager) ClientSender(msg *model.WSMessage) {
	c2gm.logger.Infof("收到訊息 處理func:SendToClient, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Msg.Data)
	conn := c2gm.GetClient(msg.ClientID)
	if conn == nil {
		c2gm.logger.Error("找不到conn")
		return
	}
	if conn.wsConn != nil {
		conn.wsConn.Write(msg.Msg)
	}
}

// GameSender 找到game 的ws 將資料ws寫給他
func (c2gm *ClientGameCommunicateManager) GameSender(msg *model.WSMessage) {
	c2gm.logger.Infof("收到訊息 處理func:SendToGame, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Msg.String())

	client := c2gm.GetClient(msg.ClientID)
	if client == nil {
		return
	}

	if !client.IsGameConnAlive() {
		c2gm.logger.Infof("gamews is not alive")
		return
	}

	if client.gameConn != nil {
		client.gameConn.Request(client.gameConn.GameWS, msg.Msg.BClassID, msg.Msg.SClassID, msg.Msg.Data)
	}
}

func (c2gm *ClientGameCommunicateManager) AddClient(conn *ClientConn) {
	c2gm.mu.Lock()
	defer c2gm.mu.Unlock()
	conn.Parser = c2gm.parser
	c2gm.client[conn.ID] = conn
}

func (c2gm *ClientGameCommunicateManager) DeleteClient(conn *ClientConn) {
	c2gm.mu.Lock()
	defer c2gm.mu.Unlock()
	delete(c2gm.client, conn.ID)
}

func (c2gm *ClientGameCommunicateManager) GetClient(clientID string) *ClientConn {
	defer c2gm.mu.Unlock()
	c2gm.mu.Lock()
	return c2gm.client[clientID]
}

func (c2gm *ClientGameCommunicateManager) GetAllClient() []*ClientConn {
	var clients = make([]*ClientConn, 0, len(c2gm.client))
	for _, c := range c2gm.client {
		clients = append(clients, c)
	}
	return clients
}
