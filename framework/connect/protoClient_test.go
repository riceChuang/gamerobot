package connect

import (
	"robot/internal/msg"
	"robot/internal/util"
	"testing"
)

func TestProtoClient(t *testing.T) {
	URL := "wss://qiangzhuangniuniu-dev.zzishare.com:5008"
	ws := NewConn(URL, nil)
	parser := &msg.ByteParser{}

	pClient := NewProtoClient(ws, parser, nil)
	pClient.Connect()

	util.WaitForCtrlC()
}

func TestProtoClientForGame(t *testing.T) {
	URL := "wss://qiangzhuangniuniu-dev.zzishare.com:7061"
	ws := NewConn(URL, nil)
	parser := &msg.ByteParser{}

	pClient := NewProtoClient(ws, parser, nil)
	pClient.Connect()

	util.WaitForCtrlC()
}
