package controller

type Controller struct {
	User *UserController
}

// All Controller
func NewController() *Controller {
	return &Controller{
		User: NewUserController(),
	}
}
