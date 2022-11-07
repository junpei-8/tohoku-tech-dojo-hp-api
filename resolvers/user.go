package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app"
	"app/ent"
	"context"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	return r.client.User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (string, error) {
	err := r.client.User.DeleteOneID(id).Exec(ctx)
	return "ok", err
}

// Mutation returns app.MutationResolver implementation.
func (r *Resolver) Mutation() app.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
