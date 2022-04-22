//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=./mock/mock_$GOFILE
package usecase

import "supermarine1377/domain"

type UserRepository interface {
	Store(u domain.User) error
	FindAll() ([]domain.User, error)
	FindById(id int) (*domain.User, error)
}
