package model

import "github.com/riceChuang/gamerobot/common"

type WSMessage struct {
	From     common.PasserType
	To       common.PasserType
	ClientID string
	Data     *Message
}

// eventChan ...
type EventChan struct {
	ReadSig  chan *Message
	WriteSig chan *Message
	ErrSig   chan error
	StopSig  chan bool
}

// ConnClient ...
type ConnClient interface {
	Connect() error
	Write(msg *Message)
	Close() error
	Clean()
	Register(h *Handler) error
	UnRegister(h *Handler)
}
