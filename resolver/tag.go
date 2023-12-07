package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/ent"
	"context"
	"log"
)

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input ent.CreateTagInput) (*ent.Tag, error) {
	log.Print("TEST")
	return r.client.Tag.Create().SetInput(input).Save(ctx)
}

// CreateTags is the resolver for the createTags field.
func (r *mutationResolver) CreateTags(ctx context.Context, inputs []*ent.CreateTagInput) ([]*ent.Tag, error) {
	tag := r.client.Tag

	bulk := make([]*ent.Tag, len(inputs))

	for i, input := range inputs {
		tag, err := tag.Create().SetInput(*input).Save(ctx)
		if err != nil {
			return nil, err
		}
		bulk[i] = tag
	}

	return bulk, nil
}

// DeleteTag is the resolver for the deleteTag field.
func (r *mutationResolver) DeleteTag(ctx context.Context, id int) (string, error) {
	err := r.client.Tag.DeleteOneID(id).Exec(ctx)
	return "ok", err
}
