package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	sqlc "github.com/Thanhbinh1905/seta-training-system/internal/db/sqlc"
)

type Resolver struct {
	Queries *sqlc.Queries
}

func NewResolver(queries *sqlc.Queries) *Resolver {
	return &Resolver{
		Queries: queries,
	}
}
