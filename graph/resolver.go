//go:generate go run github.com/99designs/gqlgen

package graph

import "github.com/axeldeveloper/go-gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	//banks []*model.Bank
}
