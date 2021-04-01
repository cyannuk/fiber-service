package repository

import (
	"fiber-service/interface/repository"
	"gopkg.in/reform.v1"

	"fiber-service/domain/model"
)

type orderRepository struct {
	DataSource
}

func (repository orderRepository) GetOrder(id int64) (*model.Order, error) {
	order, err := repository.FindByPrimaryKeyFrom(model.OrderTable, id)
	if err != nil {
		return nil, err
	}
	return order.(*model.Order), nil
}

func (repository orderRepository) GetOrders(userId int64, offset int64, limit int64) ([]model.Order, error) {
	rows, err := repository.SelectRows(model.OrderTable, "WHERE user_id = $1 ORDER BY id OFFSET $2 LIMIT $3", userId, offset, limit)
	if err != nil {
		return nil, err
	}
	orders := make([]model.Order, 0, limit)
	for {
		var order model.Order
		if err = repository.NextRow(&order, rows); err != nil {
			if err == reform.ErrNoRows {
				err = nil
			}
			_ = rows.Close()
			break
		}
		orders = append(orders, order)
	}
	return orders, err
}

func NewOrderRepository(dataSource DataSource) repository.OrderRepository {
	return orderRepository{dataSource}
}
