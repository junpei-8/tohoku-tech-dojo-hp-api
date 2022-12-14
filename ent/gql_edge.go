// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (n *News) User(ctx context.Context) (*User, error) {
	result, err := n.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = n.QueryUser().Only(ctx)
	}
	return result, err
}

func (n *News) Tags(ctx context.Context) ([]*Tag, error) {
	result, err := n.NamedTags(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = n.QueryTags().All(ctx)
	}
	return result, err
}

func (t *Tag) NewsList(ctx context.Context) ([]*News, error) {
	result, err := t.NamedNewsList(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = t.QueryNewsList().All(ctx)
	}
	return result, err
}

func (u *User) CreatedNewsList(ctx context.Context) ([]*News, error) {
	result, err := u.NamedCreatedNewsList(graphql.GetFieldContext(ctx).Field.Alias)
	if IsNotLoaded(err) {
		result, err = u.QueryCreatedNewsList().All(ctx)
	}
	return result, err
}
