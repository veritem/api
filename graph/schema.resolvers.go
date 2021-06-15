package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/veritem/api/graph/generated"
	"github.com/veritem/api/graph/model"
)

func (r *queryResolver) Name(ctx context.Context) (*model.Name, error) {
	return &model.Name{First: "Verite", Last: "Makuza", Middle: "Mugabo", Username: "veritem"}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
