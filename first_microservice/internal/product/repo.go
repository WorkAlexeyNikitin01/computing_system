package product


type ProductRepositoryInterface interface {
	AllListProducts() ([]Product, error)
	CreateProduct(p Product) (int, error)
}
