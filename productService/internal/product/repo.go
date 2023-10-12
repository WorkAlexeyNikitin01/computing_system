package product


type ProductRepositoryInterface interface {
	AllListProducts() ([]*Product, error)
	CreateProduct(p Product) (int, error)
	DeleteProduct(code string) (*Product, error)
	GetProduct(code string) (*Product, error)
}
