package order

type Repository interface {
	CreateOrder(products []interface{}) (*Order, error)
}
