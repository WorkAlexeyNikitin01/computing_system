package storeroom

type Repository interface {
	AddToStoreroom(code string, quantity int) (*StoreroomProduct, error)
	GetFromStoreroom(code string)             (*StoreroomProduct, error)
	DeleteFromStoreroom(code string)          (*StoreroomProduct, error)
}
 