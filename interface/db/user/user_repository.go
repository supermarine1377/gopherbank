package user

import (
	"supermarine1377/domain"
	"supermarine1377/interface/db"
)

type UserRepository struct {
	SqlHandler db.SqlHandler
}

func (repo UserRepository) Store(u domain.User) error {
	_, err := repo.SqlHandler.Excute("insert into users (name, balance, is_deleted) values (?, ?, ?)", u.Name, u.Balance, false)
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
	rows, err := repo.SqlHandler.Query("select id, name, balance from users where is_deleted = ?;", false)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Balance); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo UserRepository) Delete() {

}
