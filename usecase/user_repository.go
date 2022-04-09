//go:generate  -source=$GOFILE -package=mock_$GOPACKAGE -destination=./mock/$GOPACKAGE/$GOFILE
package usecase

import "supermarine1377/domain"

type UserRepository interface {
	Store(u domain.User) error
	Update(id int)
	FindAll() ([]domain.User, error)
	FindById(id int) (*domain.User, error)
}
