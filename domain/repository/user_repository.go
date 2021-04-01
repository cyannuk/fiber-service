package repository

import (
	"fiber-service/interface/repository"
	"gopkg.in/reform.v1"

	"fiber-service/domain/model"
)

type userRepository struct {
	DataSource
}

func (repository userRepository) GetUser(id int64) (*model.User, error) {
	user, err := repository.FindByPrimaryKeyFrom(model.UserTable, id)
	if err != nil {
		return nil, err
	}
	return user.(*model.User), nil
}

func (repository userRepository) DeleteUser(id int64) error {
	return repository.Delete(&model.User{ID: id})
}

func (repository userRepository) GetUsers(offset int64, limit int64) ([]model.User, error) {
	rows, err := repository.SelectRows(model.UserTable, "ORDER BY id OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, 0, limit)
	for {
		var user model.User
		if err = repository.NextRow(&user, rows); err != nil {
			if err == reform.ErrNoRows {
				err = nil
			}
			_ = rows.Close()
			break
		}
		users = append(users, user)
	}
	return users, err
}

func (repository userRepository) GetUserOrders(offset int64, limit int64) ([]model.UserOrder, error) {
	rows, err := repository.Query(
		"SELECT u.name, u.city, u.state, o.product_id, o.quantity, o.total " +
		"FROM users u " +
		"INNER JOIN orders o ON o.user_id = u.id " +
		"ORDER BY u.id, o.id " +
		"OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return nil, err
	}
	userOrders := make([]model.UserOrder, 0, limit)
	for {
		var userOrder model.UserOrder
		if err = repository.NextRow(&userOrder, rows); err != nil {
			if err == reform.ErrNoRows {
				err = nil
			}
			_ = rows.Close()
			break
		}
		userOrders = append(userOrders, userOrder)
	}
	return userOrders, nil
}

func (repository userRepository) CreateUser(user *model.User) error {
	return repository.Insert(user)
}

func (repository userRepository) FindUser(email string) (*model.User, error) {
	user := model.User{}
	return &user, repository.SelectOneTo(&user, "WHERE email = $1", email)
}

func NewUserRepository(dataSource DataSource) repository.UserRepository {
	return userRepository{dataSource}
}
