package usecase

type DsgAPIBase interface {
	//玩家登入
	Login(loginDomain string, agentID int, account string, password string, gameID int32) (url string, token string, err error)
	//入金
	StoreMoney(loginDomain string, agentID int, account string, money int32) (int32, error)
}

type AdminUseCaseBase interface {
	GetGameWsURL(hallURL string, gameRoom string, account string, token string) string
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
