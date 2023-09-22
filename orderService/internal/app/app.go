package app

import "lab/orderService/internal/order"

type AppOrderInterface interface {
	CreateOrder(products []interface{}) (*order.Order, error)
}

type AppOrderStruct struct {
	repo order.Repository
}
 
func(a *AppOrderStruct) CreateOrder(products []interface{}) (*order.Order, error) {
	return nil, nil
}

func NewAppOrder(repoOrder order.Repository) AppOrderInterface {
	return &AppOrderStruct{
		repo: repoOrder,
	}
}
