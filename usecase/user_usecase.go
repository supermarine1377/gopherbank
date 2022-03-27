package usecase

import "supermarine1377/domain"

type UserUseCase struct {
	Repository UserRepository
}

func (uu UserUseCase) Add(u domain.User) error {
	return uu.Repository.Store(u)
}

func (uu UserUseCase) FindAll() ([]domain.User, error) {
	return uu.Repository.FindAll()
}
