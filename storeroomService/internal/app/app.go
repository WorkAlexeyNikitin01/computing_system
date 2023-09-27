package app

import (
	"lab/storeroomService/internal/storeroom"
)

type AppStoreroomInterface interface {
	AddToStoreroom(sproduct storeroom.StoreroomProduct) (*storeroom.StoreroomProduct, error)
	GetFromStoreroom(code string)                       (*storeroom.StoreroomProduct, error)
	DeleteFromStoreroom(code string)                    (*storeroom.StoreroomProduct, error)
}

type AppStoreroomStruct struct {
	repo storeroom.Repository
}

func(a *AppStoreroomStruct) AddToStoreroom(sproduct storeroom.StoreroomProduct) (*storeroom.StoreroomProduct, error) {
	id, err := a.repo.AddToStoreroom(&sproduct)
	if err != nil {
		return nil, err
	}
	sproduct.Id = id
	return &sproduct, nil
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
