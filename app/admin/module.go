package admin

type Module struct {
	UserService IUserService
}

func NewModule(userService IUserService) Module {
	return Module{
		UserService: userService,
	}
}
