package usecase

import (
	"github.com/riceChuang/gamerobot/service/connect"
	"sync"
	"time"
)

var adminService *AdminService
var adminServiceOnce = &sync.Once{}

type AdminService struct {
	GameListSvc connect.GameList
}

func NewAdminService() AdminUseCaseBase {
	adminServiceOnce.Do(func() {
		adminService = &AdminService{
			GameListSvc: connect.NewGameListSvc(),
		}

	})
	return adminService
}

func (as *AdminService) GetGameWsURL(hallURL string, gameRoom string, account string, token string) string {
	var wg sync.WaitGroup
	url := as.GameListSvc.GetGameWsURL(hallURL, gameRoom, account, token)
	if url != "" {
		return url
	}
	go func() {
		wg.Add(1)
		for {
			select {
			case <-time.After(2 * time.Second):
				url = as.GameListSvc.GetGameWsURL(hallURL, gameRoom, account, token)
				if url != "" {
					wg.Done()
					return
				}
			}
		}
	}()

	wg.Wait()
	return url
}
