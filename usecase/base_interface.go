package usecase

type DsgAPIBase interface {
	//玩家登入
	Login(loginDomain string, agentID int, account string, password string, gameID int32) (string, error)
	//入金
	StoreMoney(loginDomain string, agentID int, account string, money int32) (int32, error)
}

type AdminUseCaseBase interface {
	//建立ws
	SetWebSocket()
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
