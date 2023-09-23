package app

import "lab/storeroomService/internal/storeroom"

type AppStoreroomInterface interface {
	AddToStoreroom(code string, quantity int) (*storeroom.StoreroomProduct, error)
	GetFromStoreroom(code string)             (*storeroom.StoreroomProduct, error)
	DeleteFromStoreroom(code string)          (*storeroom.StoreroomProduct, error)
}

type AppStoreroomStruct struct {
	repo storeroom.Repository
}

func(a *AppStoreroomStruct) AddToStoreroom(code string, quantity int) (*storeroom.StoreroomProduct, error) {
	return a.repo.AddToStoreroom(code, quantity)
}

func(a *AppStoreroomStruct) GetFromStoreroom(code string) (*storeroom.StoreroomProduct, error) {
	return a.repo.GetFromStoreroom(code)
}

func(a *AppStoreroomStruct) DeleteFromStoreroom(code string) (*storeroom.StoreroomProduct, error) {
	return a.repo.DeleteFromStoreroom(code)
}

func NewAppStoreroom(r storeroom.Repository) AppStoreroomInterface {
	return &AppStoreroomStruct{
		repo: r,
	}
}
