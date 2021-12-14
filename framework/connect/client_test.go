package connect

import (
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
	"github.com/riceChuang/gamerobot/util"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	URL := "wss://qiangzhuangniuniu-dev.zzishare.com:5008"
	ws := NewConn(URL)
	parser := &model.ByteParser{}

	client := NewClient(ws, parser)
	client.Connect()

	handler := &model.Handler{
		BClassID:  int32(netproto.MessageBClassID_NotifyServer),
		SClassID:  int32(netproto.NotifyServerClassID_AnnID),
		MsgType:   &netproto.AnnList{},
		OnMessage: OnAnnList,
	}

	if err := client.Register(handler); err != nil {
		log.Panicln("Rister Failed!", err.Error())
	}

	util.WaitForCtrlC()
}

func TestClientEOFAfterClean(t *testing.T) {
	URL := "wss://qiangzhuangniuniu-dev.zzishare.com:5008"
	ws := NewConn(URL)
	parser := &model.ByteParser{}

	client := NewClient(ws, parser)
	client.Connect()

	handler := &model.Handler{
		BClassID:  int32(netproto.MessageBClassID_NotifyServer),
		SClassID:  int32(netproto.NotifyServerClassID_AnnID),
		MsgType:   &netproto.AnnList{},
		OnMessage: OnAnnList,
	}

	if err := client.Register(handler); err != nil {
		log.Panicln("Rister Failed!", err.Error())
	}
	util.After(func() {
		client.Clean()
	}, 3000)

	util.WaitForCtrlC()
}

func OnAnnList(msg interface{}) {
	log.Printf("%+v", msg)
	data, ok := msg.(*netproto.AnnList)
	if !ok {
		log.Println("Error: Data Type Wrong!")
		return
	}
	log.Println("OnMessage:", data.String())
}
