// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/like"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/predicate"
)

// LikeUpdate is the builder for updating Like entities.
type LikeUpdate struct {
	config
	hooks    []Hook
	mutation *LikeMutation
}

// Where appends a list predicates to the LikeUpdate builder.
func (lu *LikeUpdate) Where(ps ...predicate.Like) *LikeUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdateTime sets the "update_time" field.
func (lu *LikeUpdate) SetUpdateTime(t time.Time) *LikeUpdate {
	lu.mutation.SetUpdateTime(t)
	return lu
}

// Mutation returns the LikeMutation object of the builder.
func (lu *LikeUpdate) Mutation() *LikeMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LikeUpdate) Save(ctx context.Context) (int, error) {
	if err := lu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LikeUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LikeUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LikeUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LikeUpdate) defaults() error {
	if _, ok := lu.mutation.UpdateTime(); !ok {
		if like.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized like.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := like.UpdateDefaultUpdateTime()
		lu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lu *LikeUpdate) check() error {
	if _, ok := lu.mutation.UserID(); lu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.user"`)
	}
	if _, ok := lu.mutation.TweetID(); lu.mutation.TweetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.tweet"`)
	}
	return nil
}

func (lu *LikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(like.FieldTweetID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdateTime(); ok {
		_spec.SetField(like.FieldUpdateTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LikeUpdateOne is the builder for updating a single Like entity.
type LikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LikeMutation
}

// SetUpdateTime sets the "update_time" field.
func (luo *LikeUpdateOne) SetUpdateTime(t time.Time) *LikeUpdateOne {
	luo.mutation.SetUpdateTime(t)
	return luo
}

// Mutation returns the LikeMutation object of the builder.
func (luo *LikeUpdateOne) Mutation() *LikeMutation {
	return luo.mutation
}

// Where appends a list predicates to the LikeUpdate builder.
func (luo *LikeUpdateOne) Where(ps ...predicate.Like) *LikeUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LikeUpdateOne) Select(field string, fields ...string) *LikeUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Like entity.
func (luo *LikeUpdateOne) Save(ctx context.Context) (*Like, error) {
	if err := luo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LikeUpdateOne) SaveX(ctx context.Context) *Like {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LikeUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LikeUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LikeUpdateOne) defaults() error {
	if _, ok := luo.mutation.UpdateTime(); !ok {
		if like.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized like.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := like.UpdateDefaultUpdateTime()
		luo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (luo *LikeUpdateOne) check() error {
	if _, ok := luo.mutation.UserID(); luo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.user"`)
	}
	if _, ok := luo.mutation.TweetID(); luo.mutation.TweetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Like.tweet"`)
	}
	return nil
}

func (luo *LikeUpdateOne) sqlSave(ctx context.Context) (_node *Like, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldUserID, field.TypeInt), sqlgraph.NewFieldSpec(like.FieldTweetID, field.TypeInt))
	if id, ok := luo.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "Like.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := luo.mutation.TweetID(); !ok {
		return nil, &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing "Like.tweet_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !like.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UpdateTime(); ok {
		_spec.SetField(like.FieldUpdateTime, field.TypeTime, value)
	}
	_node = &Like{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
