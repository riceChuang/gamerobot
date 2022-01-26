package controller

type Controller struct {
	User   *UserController
	Stress *StressController
}

// All Controller
func NewController() *Controller {
	return &Controller{
		User:   NewUserController(),
		Stress: NewStressController(),
	}
}
