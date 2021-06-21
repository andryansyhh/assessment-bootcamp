package user

type Service interface {
	GetAllUser()
	GetUserByID()
	UserRegister()
	UserLogin()
	UpdateUser()
	DeleteUser()
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}
