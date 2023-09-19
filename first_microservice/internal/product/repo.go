package product


type ProductRepositoryInterface interface {
	AllListProducts() ([]Product, error)
	CreateProduct(p Product) (Product, error)
}
