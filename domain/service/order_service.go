package service

import (
	"fiber-service/domain/model"
	"fiber-service/interface/repository"
	"fiber-service/interface/service"
)

type orderService struct {
	repository repository.OrderRepository
}

func (service orderService) GetOrder(id int64) (*model.Order, error) {
	return service.repository.GetOrder(id)
}

func (service orderService) GetOrders(userId int64, offset int64, limit int64) ([]model.Order, error) {
	return service.repository.GetOrders(userId, offset, limit)
}

func NewOrderService(repo repository.OrderRepository) service.OrderService {
	return orderService{repo}
}
