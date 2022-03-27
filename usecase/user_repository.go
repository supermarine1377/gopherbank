package usecase

import "supermarine1377/domain"

type UserRepository interface {
	Store(u domain.User) error
	Update(u domain.User)
	FindAll() ([]domain.User, error)
}
