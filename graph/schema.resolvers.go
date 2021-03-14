package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/makuzaverite/api/graph/generated"
	"github.com/makuzaverite/api/graph/model"
)

func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	return &model.User{FirstName: "Makuza", LastName: "Mugabo Verite", Age: "18", Email: "mugaboverite@gmail.com"}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
