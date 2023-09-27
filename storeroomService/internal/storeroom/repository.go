package storeroom

type Repository interface {
	AddToStoreroom(sproduct *StoreroomProduct) (int, error)
	GetFromStoreroom(code string)              (*StoreroomProduct, error)
	DeleteFromStoreroom(code string)           (*StoreroomProduct, error)
}
 