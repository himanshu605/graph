package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/himanshu605/graphql-api/graph/generated"
	"github.com/himanshu605/graphql-api/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	NUser := model.User{
		ID:   input.ID,
		Name: input.Name,
	}
	err := r.DB.Create(&NUser).Error
	if err != nil {
		return nil, err
	}
	return &NUser, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	var temp []*model.User
	r.DB.First(&temp)
	return temp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
