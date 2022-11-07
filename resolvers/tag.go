package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/ent"
	"context"
)

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input ent.CreateTagInput) (*ent.Tag, error) {
	return r.client.Tag.Create().SetInput(input).Save(ctx)
}

// DeleteTag is the resolver for the deleteTag field.
func (r *mutationResolver) DeleteTag(ctx context.Context, id int) (string, error) {
	err := r.client.Tag.DeleteOneID(id).Exec(ctx)
	return "ok", err
}
