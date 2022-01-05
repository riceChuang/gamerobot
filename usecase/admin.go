package usecase

import (
	"github.com/riceChuang/gamerobot/service/connect"
	"github.com/riceChuang/gamerobot/using/netproto"
	"sync"
	"time"
)

var adminService *AdminService
var adminServiceOnce = &sync.Once{}

type AdminService struct {
}

func NewAdminService() AdminUseCaseBase {
	adminServiceOnce.Do(func() {
		adminService = &AdminService{}
	})
	return adminService
}

func (as *AdminService) GetGameWsInfo(hallURL string, gameRoom string, account string, token string) (url string, userinfo *netproto.UserLoginRet) {
	var wg sync.WaitGroup
	hallConnect := connect.NewHallConnect()
	url, userinfo = hallConnect.GetGameWsInfo(hallURL, gameRoom, account, token)
	if url != "" && userinfo != nil {
		return
	}
	wg.Add(1)
	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				url, userinfo = hallConnect.GetGameWsInfo(hallURL, gameRoom, account, token)
				if url != "" && userinfo != nil {
					wg.Done()
					return
				}
			}
		}
	}()

	wg.Wait()
	hallConnect = nil
	return
}
