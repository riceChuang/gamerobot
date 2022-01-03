package connect

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type GameList interface {
	GetGameWsURL(hallURL string, gameRoom string, account string, token string) string
}

type GameListSvc struct {
	gameList map[string]string // map[gameid-flag] =port
	logger   *log.Entry
	Account  string
	Token    string
	hallWS   *ProtoConnect
	mu       sync.Mutex
}

func NewGameListSvc() GameList {
	return &GameListSvc{
		gameList: make(map[string]string),
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "GameListSvc",
		}),
	}
}

func (gs *GameListSvc) GetGameWsURL(hallURL string, gameRoom string, account string, token string) string {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	if url, ok := gs.gameList[gameRoom]; ok {
		return url
	} else if gs.hallWS != nil && gs.hallWS.conn.URL == hallURL {
		return ""
	} else {
		gs.Account = account
		gs.Token = token
		gs.InitHallWs(hallURL)
	}
	return ""
}

// InitHallWs ...
func (gs *GameListSvc) InitHallWs(hallURL string) {
	gs.CleanHs()
	//URL := fmt.Sprintf("%s:%d", gs.GameEnvCfg.ServerURL, gs.CommonCfg.HallPort)
	conn := framework.NewConn(hallURL, nil)
	gs.hallWS = NewProtoConnect(conn)

	gs.hallWS.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_Hall),
		SClassID:  int32(netproto.HallMsgClassID_ServerListDataID),
		MsgType:   &netproto.AllGameServerInfo{},
		OnMessage: gs.onServerListRet,
	})

	gs.hallWS.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_Hall),
		SClassID:  int32(netproto.HallMsgClassID_LoginRetID),
		MsgType:   &netproto.UserLoginRet{},
		OnMessage: gs.onLoginRet,
	})

	gs.hallWS.Connect()
	gs.requestLoginHall()
}

// CleanHs ...
func (gs *GameListSvc) CleanHs() {
	gs.logger.Infof("[DEBUG][Manager][CleanHs] ----")
	if gs.hallWS != nil {
		gs.hallWS.Clean()
		gs.hallWS = nil
	}
}

func (gs *GameListSvc) onServerListRet(msg interface{}) {
	data, ok := msg.(*netproto.AllGameServerInfo)
	if !ok {
		gs.logger.Println("Error: Data Type Wrong!")
		return
	}

	// m.logger.Log(m.Info.RobotID, "onServerListRet:", data.String())
	gs.logger.Infof("onServerListRet: %+v, user:%v", data.String(), gs.Account)
	for _, gameInfo := range data.GetServerList() {
		gameIndex := fmt.Sprintf("%v-%v", gameInfo.GetGameID(), gameInfo.GetFlag())
		gs.gameList[gameIndex] = string(gameInfo.GetPort())
	}
	go func() {
		time.Sleep(time.Second * 5)
		gs.CleanHs()
		return
	}()
	return
}

func (gs *GameListSvc) onLoginRet(msg interface{}) {
	data, ok := msg.(*netproto.UserLoginRet)
	if !ok {
		log.Println("Error: Data Type Wrong!")
		return
	}
	gs.logger.Infof("OnLoginRet: %v, user :%v", data.String(), gs.Account)
	if data.GetCode() == 1 {
		gs.RequestServerListInfo()
	} else {
		// should kill robot
		gs.logger.Errorf("未知登入錯誤: %s", data.String())
		gs.CleanHs()
	}
}

func (gs *GameListSvc) requestLoginHall() {
	gs.Request(
		gs.hallWS,
		int32(netproto.MessageBClassID_Hall),
		int32(netproto.HallMsgClassID_GuestLoginID),
		&netproto.GuestLogin{
			HDCode:     proto.String(gs.Account),
			HDType:     proto.Int32(1),
			SiteID:     proto.Int32(2),
			Version:    proto.String("123"),
			PlatformID: proto.Int32(2),
			ServerID:   proto.Int32(2),
			InviteCode: proto.String("123"),
			GuestKey:   proto.String(gs.Token),
		},
	)
}

// RequestServerListInfo ...
func (gs *GameListSvc) RequestServerListInfo() {
	gs.Request(
		gs.hallWS,
		int32(netproto.MessageBClassID_Hall),
		int32(netproto.HallMsgClassID_RequestServerListID),
		&netproto.GetShowGameFlagByAreaMsg{
			Area_ID: proto.Int32(2),
		},
	)
}

// Request send and handle protoMsg Error
func (gs *GameListSvc) Request(ws *ProtoConnect, bclsID int32, sclsID int32, protoData interface{}) {
	message, err := model.NewProto(bclsID, sclsID, protoData)

	// should kill robot
	if err != nil {
		gs.logger.Panic(err.Error())
	}
	ws.Write(message)
}
