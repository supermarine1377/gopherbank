package user

import (
	"fmt"
	"supermarine1377/domain"
	"supermarine1377/interface/db"
)

type UserRepository struct {
	SqlHandler db.SqlHandler
}

func (repo UserRepository) Store(u domain.User) error {
	statement := fmt.Sprintf(`insert into users (name, balance, is_deleted) values ('%s', %d, false);`, u.Name, u.Balance)
	_, err := repo.SqlHandler.Excute(statement)
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepository) Update(u domain.User) {

}

func (repo UserRepository) FindById(id int) {

}

func (repo UserRepository) FindAll() ([]domain.User, error) {
	statement := "select id, name, balance from users where is_deleted = false"
	row, err := repo.SqlHandler.Query(statement)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for row.Next() {
		var id int
		var name string
		var balance int
		if err := row.Scan(&id, &name, &balance); err != nil {
			continue
		}
		users = append(users, domain.User{
			ID:      id,
			Name:    name,
			Balance: balance,
		})
	}
	return users, nil
}

func (repo UserRepository) Delete() {

}
