package app

import "lab/orderService/internal/order"

type AppOrderInterface interface {
	AddToOrder() (*order.Order, error)
}

type AppOrderStruct struct {
	repo order.Repository
}
 
func(a *AppOrderStruct) AddToOrder() (*order.Order, error) {
	return nil, nil
}

func NewAppOrder(repoOrder order.Repository) AppOrderInterface {
	return &AppOrderStruct{
		repo: repoOrder,
	}
}
