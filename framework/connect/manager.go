package connect

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/service/usecase"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"
	"net/url"
)

const (
	areaID int32 = 2
)

// Manager ...
type Manager struct {
	Account  string
	Passwrod string

	UserMoney int32

	userInfo    *netproto.UserLoginRet
	gameServers *netproto.AllGameServerInfo

	// by generate
	Token       string
	GameEnvCfg  *model.EnvCfg
	CommonCfg   *model.CommonCfg
	RoomCfgBase *model.RoomCfgBase

	logger *log.Entry
	hallWs ConnClient
	gameWs ConnClient
}

// NewManager ...
func NewManager(gameEnv *model.CommonCfg, env *model.EnvCfg) (*Manager, error) {
	m := &Manager{}
	m.GameEnvCfg = env
	m.CommonCfg = gameEnv
	m.logger = logs.GetLogger().WithFields(
		log.Fields{
			"server": "manager",
		})
	err := m.initToken(m.Account, m.CommonCfg.GameID)
	return m, err
}

// InitHallWs ...
func (m *Manager) InitHallWs() {
	m.CleanHs()
	URL := fmt.Sprintf("%s:%d", m.GameEnvCfg.ServerURL, m.CommonCfg.HallPort)
	conn := NewConn(URL)
	parser := &model.ByteParser{}
	m.hallWs = NewProtoClient(conn, parser)

	m.hallWs.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_Hall),
		SClassID:  int32(netproto.HallMsgClassID_ServerListDataID),
		MsgType:   &netproto.AllGameServerInfo{},
		OnMessage: m.onServerListRet,
	})

	m.hallWs.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_Hall),
		SClassID:  int32(netproto.HallMsgClassID_LoginRetID),
		MsgType:   &netproto.UserLoginRet{},
		OnMessage: m.onLoginRet,
	})

	m.hallWs.Connect()
	m.requestLoginHall()
}

// InitGameWs ...
func (m *Manager) InitGameWs(gamePort int32) {
	m.CleanGs()
	URL := fmt.Sprintf("%s:%d", m.GameEnvCfg.ServerURL, gamePort)
	conn := NewConn(URL)
	parser := &model.ByteParser{}
	m.gameWs = NewProtoClient(conn, parser)

	m.gameWs.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_GameRoom),
		SClassID:  int32(netproto.GameRoomClassID_LoginRoomRetID),
		MsgType:   &netproto.LoginGameRoomRet{},
		OnMessage: m.onGameLogin,
	})

	m.gameWs.Connect()
	m.requestLoginGame()
}

// StoreMoney ...
func (m *Manager) StoreMoney() {
	APIDomain := m.GameEnvCfg.LoginDomain
	agentID := m.GameEnvCfg.AgentID
	account := m.Account
	money := m.RoomCfgBase.StoredMoney

	_, err := usecase.NewDSGService().StoreMoney(APIDomain, agentID, account, money)
	if err != nil {
		m.logger.Error(err.Error())
	}
}

// CleanHs ...
func (m *Manager) CleanHs() {
	m.logger.Debug("[DEBUG][Manager][CleanHs] ----")
	if m.hallWs != nil {
		m.hallWs.Clean()
		m.hallWs = nil
	}
}

// CleanGs ...
func (m *Manager) CleanGs() {
	m.logger.Debug("[DEBUG][Manager][CleanGs] ----")
	if m.gameWs != nil {
		m.gameWs.Clean()
		m.gameWs = nil
	}
}

// GetGameServers ...
func (m *Manager) GetGameServers() *netproto.AllGameServerInfo {
	return m.gameServers
}

// GetUserInfo ...
func (m *Manager) GetUserInfo() *netproto.UserLoginRet {
	return m.userInfo
}

// GetUserID ...
func (m *Manager) GetUserID() int32 {
	if m.userInfo != nil {
		return m.userInfo.GetUserID()
	}
	return 0
}

// GetIsGaming ...
func (m *Manager) GetIsGaming() bool {
	return m.gameWs != nil
}

// GetGameWs ...
func (m *Manager) GetGameWs() ConnClient {
	return m.gameWs
}


// ---- Hall ----

func (m *Manager) onLoginRet(msg interface{}) {
	data, ok := msg.(*netproto.UserLoginRet)
	if !ok {
		log.Println("Error: Data Type Wrong!")
		return
	}
	m.logger.Infof( "OnLoginRet:%v", data.String())
	if data.GetCode() == 1 {
		m.userInfo = data
		m.UserMoney = int32(m.userInfo.GetUserData().GetCashAmount() / int64(100))
		m.RequestServerListInfo()
	} else {
		// should kill robot
		m.logger.Panic(fmt.Errorf("未知登入錯誤: %s", data.String()))
	}
}

func (m *Manager) onServerListRet(msg interface{}) {
	data, ok := msg.(*netproto.AllGameServerInfo)
	if !ok {
		log.Println("Error: Data Type Wrong!")
		return
	}

	// m.logger.Log(m.Info.RobotID, "onServerListRet:", data.String())
	m.logger.Infof("onServerListRet: %+v", data.String())
	m.gameServers = data
}

func (m *Manager) requestLoginHall() {
	m.Request(
		m.hallWs,
		int32(netproto.MessageBClassID_Hall),
		int32(netproto.HallMsgClassID_GuestLoginID),
		&netproto.GuestLogin{
			HDCode:     proto.String(m.Account),
			HDType:     proto.Int32(1),
			SiteID:     proto.Int32(2),
			Version:    proto.String("123"),
			PlatformID: proto.Int32(2),
			ServerID:   proto.Int32(2),
			InviteCode: proto.String("123"),
			GuestKey:   proto.String(m.Token),
		},
	)
}

// RequestServerListInfo ...
func (m *Manager) RequestServerListInfo() {
	m.Request(
		m.hallWs,
		int32(netproto.MessageBClassID_Hall),
		int32(netproto.HallMsgClassID_RequestServerListID),
		&netproto.GetShowGameFlagByAreaMsg{
			Area_ID: proto.Int32(areaID),
		},
	)
}

// ---- Game ----

func (m *Manager) requestLoginGame() {
	m.Request(
		m.gameWs,
		int32(netproto.MessageBClassID_GameRoom),
		int32(netproto.GameRoomClassID_LoginRoomID),
		&netproto.LoginGameRoomInfo{
			UserID: m.userInfo.UserID,
			Cer:    m.userInfo.Cer,
			HDCode: m.userInfo.HDCode,
			HDType: m.userInfo.HDType,
		},
	)
}

func (m *Manager) onGameLogin(msg interface{}) {
	data, ok := msg.(*netproto.LoginGameRoomRet)
	if !ok {
		log.Println("Error: Data Type Wrong!")
		return
	}
	m.logger.Infof("onGameLogin:%+v", data.String())
}

// ---- Common ----

// initToken if not get will throw panic make robot shut down
func (m *Manager) initToken(account string, gameID int32) error {
	APIDomain := m.GameEnvCfg.LoginDomain
	agentID := m.GameEnvCfg.AgentID
	URL, err := usecase.NewDSGService().Login(APIDomain, agentID, account, gameID)
	if err != nil {
		log.Fatalf("Error: Login error: %v", err)
		return err
	}

	u, err := url.Parse(URL)
	if err != nil {
		log.Fatalf("Error: Login parse error: %v", err)
		return err
	}

	query := u.Query()
	if token, ok := query["token"]; ok {
		m.Token = token[0]
		//m.logger.Infof("%d success get token %s", m.Info.RobotID, m.Token)

	} else {
		return fmt.Errorf("login success but no token %v", query)
	}
	return nil
}

// Request send and handle protoMsg Error
func (m *Manager) Request(ws ConnClient, bclsID int32, sclsID int32, protoData interface{}) {
	message, err := model.NewProto(bclsID, sclsID, protoData)

	// should kill robot
	if err != nil {
		m.logger.Panic(err.Error())
	}
	ws.Write(message)
}

//func (m *Manager) GetRoomConfig() *map[string]interface{} {
//	return m.cfgLoader.GetDynamicRobotRoomConfig(m.Info.GameID, m.Info.ConfigID, m.Info.RoomIndex)
//}
