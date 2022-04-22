//go:generate mockgen -source=$GOFILE -package=mock -destination=./mock/mock_$GOFILE
package infrastructure

import "supermarine1377/domain"

type UserController interface {
	Add(user domain.User) error
	FindAll() ([]domain.User, error)
	FindById(id int) (*domain.User, error)
}
