package connect

import (
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util"
	"time"
)

// ProtoClient wrapper client
type ProtoClient struct {
	Client
	timerChan    chan bool
	ClientErrSig chan error
}

// NewProtoClient new a net client
func NewProtoClient(conn *Conn, parser model.Parser) *ProtoClient {
	pc := &ProtoClient{}
	pc.Client = *NewClient(conn, parser)
	pc.onCloseDelegate = pc.onClose
	pc.ClientErrSig = make(chan error)
	return pc
}

// Connect @override add heartbeat setup
func (pc *ProtoClient) Connect() error {
	err := pc.Client.Connect()
	if err != nil {
		return err
	}
	pc.initHeartBeat()
	return nil
}

// onClose delegate
func (pc *ProtoClient) onClose() {
	pc.Client.logger.Printf("ProtoClient onClose")
	if pc.timerChan != nil {
		pc.timerChan <- true
		pc.timerChan = nil
	}
}

func (pc *ProtoClient) initHeartBeat() {
	pc.Client.Register(&model.Handler{
		BClassID:  int32(netproto.MessageBClassID_Common),
		SClassID:  int32(netproto.PlatformCommonClassID_HeartBeatConfigID),
		MsgType:   &netproto.HeartBeatConfig{},
		OnMessage: pc.onHeartBeatCfg,
	})

	pc.requestHeartBeatCfg()
}

func (pc *ProtoClient) requestHeartBeatCfg() {
	message, err := model.NewProto(
		int32(netproto.MessageBClassID_Common),
		int32(netproto.PlatformCommonClassID_RequestHeartBeatConfigID),
		nil,
	)

	if err != nil {
		pc.Client.logger.Println("Error: send heart beat failed", err.Error())
	}
	pc.Client.Write(message)
}

func (pc *ProtoClient) onHeartBeatCfg(b interface{}) {
	data, ok := b.(*netproto.HeartBeatConfig)
	if !ok {
		// should kill robot
		pc.Client.logger.Panic("Error: Data Type Wrong!")
		return
	}
	pc.Client.logger.Println("onHeartBeatCfg:", data.String())

	if pc.timerChan != nil {
		pc.timerChan <- true
	}
	pc.timerChan = util.NewTicker(pc.sendHeartBeat, time.Duration(data.GetBeatInterval())*time.Second)
}

func (pc *ProtoClient) sendHeartBeat() {
	message, err := model.NewProto(
		int32(netproto.MessageBClassID_Common),
		int32(netproto.PlatformCommonClassID_HeartBeatID),
		nil,
	)

	if err != nil {
		pc.Client.logger.Println("Error: send heart beat failed", err.Error())
	}
	pc.Client.Write(message)
}
