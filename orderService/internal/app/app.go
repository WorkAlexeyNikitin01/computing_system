package app

import "lab/orderService/internal/order"

type AppOrderInterface interface {
	CreateOrder(order order.Order) (*order.Order, error)
}

type AppOrderStruct struct {
	repo order.Repository
}
 
func(a *AppOrderStruct) CreateOrder(order order.Order) (*order.Order, error) {
	id, err := a.repo.CreateOrder(&order)
	if err != nil {
		return nil, err
	}
	order.Id = id
	return &order, nil
}

func NewAppOrder(repoOrder order.Repository) AppOrderInterface {
	return &AppOrderStruct{
		repo: repoOrder,
	}
}
