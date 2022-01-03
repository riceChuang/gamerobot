package connect

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/service/game"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
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
	parser       framework.Parser
	gameWSUrlMap map[string]string
	hallWS       *ProtoConnect
}

func NewClientWsToGameServer(wsAddr string) *ClientGameCommunicateManager {
	clientGameCommunicateOnce.Do(func() {
		clientGameCommunicateManager = &ClientGameCommunicateManager{
			mu:          &sync.Mutex{},
			wsSeverAddr: wsAddr,
			client:      make(map[string]*ClientConn),
			logger: logs.GetLogger().WithFields(log.Fields{
				"server": "websocket",
			}),
			parser: &framework.ByteParser{},
		}

		framework.GetClientDispatcher().AddMsgPasser(string(common.ServerToTransfer), clientGameCommunicateManager.TransferClientMessage)
		framework.GetClientDispatcher().AddMsgPasser(string(common.TransferToClient), clientGameCommunicateManager.SendToClient)
		framework.GetGameDispatcher().AddMsgPasser(string(common.ServerToTransfer), clientGameCommunicateManager.TransferGameMessage)
		framework.GetGameDispatcher().AddMsgPasser(string(common.TransferToGame), clientGameCommunicateManager.SendToGame)
	})
	return clientGameCommunicateManager
}

func GetClientGameCommunicateManager() *ClientGameCommunicateManager {
	return clientGameCommunicateManager
}

// TransferClientMessage 轉換從client端來資料,處理後丟給遊戲
func (c2gm *ClientGameCommunicateManager) TransferClientMessage(msg *model.WSMessage) {
	//轉換或處理資料
	/*
		1. 依照看是哪個game的資訊找尋相對應的game_factory 將訊息丟給工廠處理
		2. 工廠若處理完畢會再送往傳送game的通道
	*/
	c2gm.logger.Infof("收到訊息 處理func:TransferClientMessage, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Data.String())
	msg.From = common.ServerToTransfer
	msg.To = common.TransferToGame
	framework.GetGameDispatcher().AddMessage(msg)
}

// TransferGameMessage 轉換game端來的資料,處理後丟給client
func (c2gm *ClientGameCommunicateManager) TransferGameMessage(msg *model.WSMessage) {
	//轉換或處理資料
	/*
		1. 需要把拿game的資訊丟一份往client顯示
		2. 依照看是哪個game的資訊找尋相對應的game_factory 將訊息丟給工廠處理
		3. 工廠處理完訊息可能會有部分訊息網game通道送 有可能也會網client送
	*/
	c2gm.logger.Infof("收到訊息 處理func:TransferGameMessage, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Data.String())
	client := c2gm.GetClient(msg.ClientID)
	gameInstance := game.GetInstanceByContentType(client.GameID)
	if gameInstance == nil {
		c2gm.logger.Errorf("找不到game instance %v", client.GameID)
	}

	gameInstance.HandleMessage()

	msg.From = common.ServerToTransfer
	msg.To = common.TransferToClient
	framework.GetClientDispatcher().AddMessage(msg)
}

// SendToClient 找到client 的ws 將資料ws寫給他
func (c2gm *ClientGameCommunicateManager) SendToClient(msg *model.WSMessage) {
	c2gm.logger.Infof("收到訊息 處理func:SendToClient, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Data.String())
	conn := c2gm.GetClient(msg.ClientID)
	if conn == nil {
		c2gm.logger.Error("找不到conn")
		return
	}
	conn.wsConn.Write(msg.Data)
}

// SendToGame 找到game 的ws 將資料ws寫給他
func (c2gm *ClientGameCommunicateManager) SendToGame(msg *model.WSMessage) {
	c2gm.logger.Infof("收到訊息 處理func:SendToGame, from:%v, to:%v, 資訊內容:%v", msg.From, msg.To, msg.Data.String())
	msg.From = common.GameToServer
	msg.To = common.ServerToTransfer
	framework.GetGameDispatcher().AddMessage(msg)
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

func (c2gm *ClientGameCommunicateManager) CheckTcpAlive() bool {
	c2gm.mu.Lock()
	defer c2gm.mu.Unlock()

	//conn, err := net.Dial("tcp", c2gm.tcpServerAddr)
	//if err != nil {
	//	c2gm.logger.Error(("httpsrc other tcp[%v] server error", c2gm.tcpServerAddr)
	//	return false
	//}

	//conn.Close()
	return true
}

