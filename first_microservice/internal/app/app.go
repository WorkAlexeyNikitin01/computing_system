package app

import product "lab/first_microservice/internal/product"

type AppProductInterface interface{
	AllListProducts() ([]product.Product, error)
	CreateProduct(name string) (*product.Product, error)
}

type AppProductStruct struct {
	repo product.ProductRepositoryInterface
}

func(a *AppProductStruct) AllListProducts() ([]product.Product, error) {
	listProducts, err := a.repo.AllListProducts()
	return listProducts, err
}

func(a *AppProductStruct) CreateProduct(name string) (*product.Product, error) {
	return nil, nil
}

func NewAppProduct(repoProduct product.ProductRepositoryInterface) AppProductInterface {
	return &AppProductStruct{
		repo: repoProduct,
	}
}
