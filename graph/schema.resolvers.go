package graph

import (
	"context"
	"fmt"

	"github.com/makuzaverite/api/graph/generated"
	"github.com/makuzaverite/api/graph/model"
)

func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	var user *model.User

	user.FirstName = "Makuza"
	user.LastName = "Mugabo"
	user.Age = "18"
	user.Email = "mugaboverite@gmail.com"

	return &model.User{}, fmt.Errorf("No use found")
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
