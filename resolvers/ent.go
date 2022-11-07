package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app"
	"app/ent"
	"context"
	"fmt"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// NewsList is the resolver for the news_list field.
func (r *queryResolver) NewsList(ctx context.Context) ([]*ent.News, error) {
	panic(fmt.Errorf("not implemented: NewsList - news_list"))
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context) ([]*ent.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Query returns app.QueryResolver implementation.
func (r *Resolver) Query() app.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
