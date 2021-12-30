package usecase

import "sync"

var adminService *AdminService
var adminServiceOnce = &sync.Once{}

type AdminService struct{}

func NewAdminService() AdminUseCaseBase {
	dsgApiServiceOnce.Do(func() {
		adminService = &AdminService{}
	})
	return adminService
}

func (as *AdminService) SetWebSocket() {

}
