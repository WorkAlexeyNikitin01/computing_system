package graph

import "lab/apiGraphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func(r *queryResolver) GetProduct(code string) (*model.Product) {
	return &model.Product{
		ID: "10",
		Code: "code",
		Name: "name",
		Price: 100,
	}
} 
