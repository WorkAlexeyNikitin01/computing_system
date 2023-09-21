package order

type Repository interface {
	AddToOrder() (*Order, error)
}