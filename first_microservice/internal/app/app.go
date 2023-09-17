package app

import (
	"fmt"
)

func printApp() {
	fmt.Println("hello App first service")
}

type AppOrder struct{}

type AppOrderInterface interface{}

func NewAppOrder() {}
