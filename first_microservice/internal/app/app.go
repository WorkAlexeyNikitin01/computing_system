package app

import product "lab/first_microservice/internal/product"

type AppProductInterface interface{
	AllListProducts() ([]product.Product, error)
	CreateProduct(p product.Product) (*product.Product, error)
}

type AppProductStruct struct {
	repo product.ProductRepositoryInterface
}

func(a *AppProductStruct) AllListProducts() ([]product.Product, error) {
	listProducts, err := a.repo.AllListProducts()
	return listProducts, err
}

func(a *AppProductStruct) CreateProduct(p product.Product) (*product.Product, error) {
	id, err := a.repo.CreateProduct(p)
	if err != nil {
		return nil, err
	}
	p.Id = id
	return &p, nil
}

func NewAppProduct(repoProduct product.ProductRepositoryInterface) AppProductInterface {
	return &AppProductStruct{
		repo: repoProduct,
	}
}
