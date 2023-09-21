package storeroom

type Repository interface {
	AddToStoreroom()                     (interface{}, error)
	GetProductFromStoreroom(code string) (interface{}, error)
}
