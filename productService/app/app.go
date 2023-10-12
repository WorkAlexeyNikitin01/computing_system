package app

import product "lab/productService/product"

type AppProductInterface interface{
	AllListProducts() ([]*product.Product, error)
	CreateProduct(p product.Product) (*product.Product, error)
	DeleteProduct(code string) (*product.Product, error)
	GetProduct(code string) (*product.Product, error)
}

type AppProductStruct struct {
	Repo product.ProductRepositoryInterface
}

func(a *AppProductStruct) AllListProducts() ([]*product.Product, error) {
	listProducts, err := a.Repo.AllListProducts()
	return listProducts, err
}

func(a *AppProductStruct) CreateProduct(p product.Product) (*product.Product, error) {
	id, err := a.Repo.CreateProduct(p)
	if err != nil {
		return nil, err
	}
	p.Id = id
	return &p, nil
}

func(a *AppProductStruct) DeleteProduct(code string) (*product.Product, error) {
	return a.Repo.DeleteProduct(code)
}

func(a *AppProductStruct) GetProduct(code string) (*product.Product, error) {
	p, err := a.Repo.GetProduct(code)
	return p, err
}

func NewAppProduct(repoProduct product.ProductRepositoryInterface) AppProductInterface {
	return &AppProductStruct{
		Repo: repoProduct,
	}
}
