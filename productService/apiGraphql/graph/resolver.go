package graph

import "lab/productService/internal/app"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AppProduct app.AppProductInterface
}
