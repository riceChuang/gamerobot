package framework

import (
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/logs"
	"sync"
)

//用來做接收發的處理器
type Dispatcher interface {
	AddMsgPasser(name string, passerFunc func(data *model.WSMessage)) bool
	AddMessage(data *model.WSMessage)
}

type Dispatch struct {
	Name         string
	receiveQueue chan *model.WSMessage
	passerMap    map[string]func(data *model.WSMessage)
	stopSig      chan bool
	mu           sync.Mutex
}


var ClientDispatcher Dispatcher
var clientDispatcherOnce sync.Once
var GameDispatcher Dispatcher
var gameDispatcherOnce sync.Once

func InitDispatcher() {
	clientDispatcherOnce.Do(func() {
		ClientDispatcher = NewDispatcher(string(common.ClientWSDispatcher))
	})
	gameDispatcherOnce.Do(func() {
		GameDispatcher = NewDispatcher(string(common.GameWSDispatcher))
	})
}

func GetClientDispatcher() Dispatcher {
	return ClientDispatcher
}

func GetGameDispatcher() Dispatcher {
	return GameDispatcher
}

func NewDispatcher(name string) Dispatcher {
	dp := &Dispatch{
		Name:         name,
		receiveQueue: make(chan *model.WSMessage, 1000),
		stopSig:      make(chan bool, 1),
		passerMap:    make(map[string]func(data *model.WSMessage)),
	}
	go dp.handleMsg()
	return dp
}

//add passer
func (dp *Dispatch) AddMsgPasser(name string, passerFunc func(data *model.WSMessage)) bool {
	dp.mu.Lock()
	dp.passerMap[name] = passerFunc
	dp.mu.Unlock()
	return true
}

func (dp *Dispatch) AddMessage(data *model.WSMessage) {
	dp.receiveQueue <- data
}

//delete passer
func (dp *Dispatch) DeleteMsgPasser(name string) bool {
	dp.mu.Lock()
	delete(dp.passerMap, name)
	dp.mu.Unlock()
	return true
}

func (dp *Dispatch) handleMsg() {
	defer func() {
		if r := recover(); r != nil {
			logs.GetLogger().Error("handleMsg error:%v", r.(error))
			go dp.handleMsg()
			return
		}
	}()

HandleLoop:
	for {
		select {
		case message := <-dp.receiveQueue:
			dp.handleReceiveQueue(message)
		case <-dp.stopSig:
			break HandleLoop
		}
	}
}

func (dp *Dispatch) Release() {
	dp.stopSig <- true
	return
}

func (dp *Dispatch) handleReceiveQueue(data *model.WSMessage) {
	if f, ok := dp.passerMap[string(data.To)]; ok {
		f(data)
	} else {
		logs.GetLogger().Error("找不到passer %v", data)
	}
}
