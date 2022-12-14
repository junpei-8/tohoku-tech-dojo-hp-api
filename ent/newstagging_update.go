// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/newstagging"
	"app/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsTaggingUpdate is the builder for updating NewsTagging entities.
type NewsTaggingUpdate struct {
	config
	hooks    []Hook
	mutation *NewsTaggingMutation
}

// Where appends a list predicates to the NewsTaggingUpdate builder.
func (ntu *NewsTaggingUpdate) Where(ps ...predicate.NewsTagging) *NewsTaggingUpdate {
	ntu.mutation.Where(ps...)
	return ntu
}

// Mutation returns the NewsTaggingMutation object of the builder.
func (ntu *NewsTaggingUpdate) Mutation() *NewsTaggingMutation {
	return ntu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ntu *NewsTaggingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ntu.hooks) == 0 {
		if err = ntu.check(); err != nil {
			return 0, err
		}
		affected, err = ntu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsTaggingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ntu.check(); err != nil {
				return 0, err
			}
			ntu.mutation = mutation
			affected, err = ntu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ntu.hooks) - 1; i >= 0; i-- {
			if ntu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ntu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ntu *NewsTaggingUpdate) SaveX(ctx context.Context) int {
	affected, err := ntu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ntu *NewsTaggingUpdate) Exec(ctx context.Context) error {
	_, err := ntu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntu *NewsTaggingUpdate) ExecX(ctx context.Context) {
	if err := ntu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntu *NewsTaggingUpdate) check() error {
	if _, ok := ntu.mutation.NewsID(); ntu.mutation.NewsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NewsTagging.news"`)
	}
	if _, ok := ntu.mutation.TagID(); ntu.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NewsTagging.tag"`)
	}
	return nil
}

func (ntu *NewsTaggingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newstagging.Table,
			Columns: newstagging.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: newstagging.FieldNewsID,
				},
				{
					Type:   field.TypeInt,
					Column: newstagging.FieldTagID,
				},
			},
		},
	}
	if ps := ntu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ntu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newstagging.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// NewsTaggingUpdateOne is the builder for updating a single NewsTagging entity.
type NewsTaggingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewsTaggingMutation
}

// Mutation returns the NewsTaggingMutation object of the builder.
func (ntuo *NewsTaggingUpdateOne) Mutation() *NewsTaggingMutation {
	return ntuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ntuo *NewsTaggingUpdateOne) Select(field string, fields ...string) *NewsTaggingUpdateOne {
	ntuo.fields = append([]string{field}, fields...)
	return ntuo
}

// Save executes the query and returns the updated NewsTagging entity.
func (ntuo *NewsTaggingUpdateOne) Save(ctx context.Context) (*NewsTagging, error) {
	var (
		err  error
		node *NewsTagging
	)
	if len(ntuo.hooks) == 0 {
		if err = ntuo.check(); err != nil {
			return nil, err
		}
		node, err = ntuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsTaggingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ntuo.check(); err != nil {
				return nil, err
			}
			ntuo.mutation = mutation
			node, err = ntuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ntuo.hooks) - 1; i >= 0; i-- {
			if ntuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ntuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*NewsTagging)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NewsTaggingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ntuo *NewsTaggingUpdateOne) SaveX(ctx context.Context) *NewsTagging {
	node, err := ntuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ntuo *NewsTaggingUpdateOne) Exec(ctx context.Context) error {
	_, err := ntuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntuo *NewsTaggingUpdateOne) ExecX(ctx context.Context) {
	if err := ntuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntuo *NewsTaggingUpdateOne) check() error {
	if _, ok := ntuo.mutation.NewsID(); ntuo.mutation.NewsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NewsTagging.news"`)
	}
	if _, ok := ntuo.mutation.TagID(); ntuo.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NewsTagging.tag"`)
	}
	return nil
}

func (ntuo *NewsTaggingUpdateOne) sqlSave(ctx context.Context) (_node *NewsTagging, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   newstagging.Table,
			Columns: newstagging.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: newstagging.FieldNewsID,
				},
				{
					Type:   field.TypeInt,
					Column: newstagging.FieldTagID,
				},
			},
		},
	}
	if id, ok := ntuo.mutation.NewsID(); !ok {
		return nil, &ValidationError{Name: "news_id", err: errors.New(`ent: missing "NewsTagging.news_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ntuo.mutation.TagID(); !ok {
		return nil, &ValidationError{Name: "tag_id", err: errors.New(`ent: missing "NewsTagging.tag_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ntuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !newstagging.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ntuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &NewsTagging{config: ntuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ntuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{newstagging.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
