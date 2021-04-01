package service

import (
	"fiber-service/domain/model"
)

type OrderService interface {
	GetOrder(id int64) (*model.Order, error)
	GetOrders(userId int64, offset int64, limit int64) (orders []model.Order, err error)
}
