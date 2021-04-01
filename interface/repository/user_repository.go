package repository

import (
	"fiber-service/domain/model"
)

type UserRepository interface {
	GetUser(id int64) (*model.User, error)
	DeleteUser(id int64) error
	GetUsers(offset int64, limit int64) (users []model.User, err error)
	GetUserOrders(offset int64, limit int64) (userOrders []model.UserOrder, err error)
	CreateUser(user *model.User) error
	FindUser(email string) (*model.User, error)
}
