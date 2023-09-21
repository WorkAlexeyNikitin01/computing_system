package app

import "lab/storeroomService/internal/storeroom"

type AppStoreroomInterface interface {
	AddToStoreroom() (interface{}, error)
	GetProductFromStoreroom(code string) (interface{}, error)
}

type AppStoreroomStruct struct {
	repo storeroom.Repository
}

func(a *AppStoreroomStruct) AddToStoreroom() (interface{}, error) {
	return nil, nil
}

func(a *AppStoreroomStruct) GetProductFromStoreroom(code string) (interface{}, error) {
	return nil, nil
}

func NewAppStoreroom(r storeroom.Repository) AppStoreroomInterface {
	return &AppStoreroomStruct{
		repo: r,
	}
}
