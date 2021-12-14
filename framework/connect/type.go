package connect

import "github.com/riceChuang/gamerobot/model"

// eventChan ...
type eventChan struct {
	ReadSig  chan *model.Message
	WriteSig chan *model.Message
	ErrSig   chan error
	StopSig  chan bool
}

// ConnClient ...
type ConnClient interface {
	Connect() error
	Write(msg *model.Message)
	Close() error
	Clean()
	Register(h *model.Handler) error
	UnRegister(h *model.Handler)
}
