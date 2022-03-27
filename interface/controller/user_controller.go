package controller

import (
	"supermarine1377/domain"
	"supermarine1377/interface/db"
	"supermarine1377/interface/db/user"
	"supermarine1377/usecase"
)

type UserController struct {
	usecase usecase.UserUseCase
}

func NewUserController(sqlHandler db.SqlHandler) *UserController {
	return &UserController{
		usecase: usecase.UserUseCase{
			Repository: user.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller UserController) Add(user domain.User) error {
	return controller.usecase.Add(user)
}

func (controller UserController) FindAll() ([]domain.User, error) {
	return controller.usecase.FindAll()
}
