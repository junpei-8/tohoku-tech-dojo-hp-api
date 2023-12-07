package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app"
	"app/ent"
	"context"
)

// CreateNews is the resolver for the createNews field.
func (r *mutationResolver) CreateNews(ctx context.Context, input ent.CreateNewsInput) (*ent.News, error) {
	return r.client.News.Create().SetInput(input).Save(ctx)
}

// UpdateNews is the resolver for the updateNews field.
func (r *mutationResolver) UpdateNews(ctx context.Context, id int, input ent.UpdateNewsInput) (*ent.News, error) {
	return r.client.News.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteNews is the resolver for the deleteNews field.
func (r *mutationResolver) DeleteNews(ctx context.Context, id int) (string, error) {
	err := r.client.News.DeleteOneID(id).Exec(ctx)
	return "ok", err
}

// Mutation returns app.MutationResolver implementation.
func (r *Resolver) Mutation() app.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
