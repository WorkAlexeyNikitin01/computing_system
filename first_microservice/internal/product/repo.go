package product


type ProductRepositoryInterface interface {
	AllListProducts() ([]Product, error)
	CreateProduct(name string) (Product, error)
}
