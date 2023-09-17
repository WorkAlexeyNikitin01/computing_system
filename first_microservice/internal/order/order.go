package order

import (
	"fmt"
)

func printFields() {
	fmt.Println("model for order")
}

type Order struct {
	Price       float64
	Quantity    float64
	Description string
}

type OrderInterface interface {
	Pay(q float64) (float64, error)
}
