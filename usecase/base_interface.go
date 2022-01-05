package usecase

import (
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/using/netproto"
)

type DsgAPIBase interface {
	//玩家登入
	Login(loginReq *model.DSGLoginReq) (loginResp *model.DSGLoginResp, err error)
	//入金
	StoreMoney(loginDomain string, agentID int, account string, money int32) (int32, error)
}

type AdminUseCaseBase interface {
	GetGameWsInfo(hallURL string, gameRoom string, account string, token string) (string, *netproto.UserLoginRet)
}

type UseCase struct {
	DsgAPIBase DsgAPIBase
	AdminBase  AdminUseCaseBase
}

func NewUseCase() *UseCase {
	useCase := &UseCase{
		DsgAPIBase: NewDSGService(),
		AdminBase:  NewAdminService(),
	}
	return useCase
}
