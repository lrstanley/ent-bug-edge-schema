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
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/tweet"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/user"
)

// LikeCreate is the builder for creating a Like entity.
type LikeCreate struct {
	config
	mutation *LikeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (lc *LikeCreate) SetCreateTime(t time.Time) *LikeCreate {
	lc.mutation.SetCreateTime(t)
	return lc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (lc *LikeCreate) SetNillableCreateTime(t *time.Time) *LikeCreate {
	if t != nil {
		lc.SetCreateTime(*t)
	}
	return lc
}

// SetUpdateTime sets the "update_time" field.
func (lc *LikeCreate) SetUpdateTime(t time.Time) *LikeCreate {
	lc.mutation.SetUpdateTime(t)
	return lc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (lc *LikeCreate) SetNillableUpdateTime(t *time.Time) *LikeCreate {
	if t != nil {
		lc.SetUpdateTime(*t)
	}
	return lc
}

// SetLikedAt sets the "liked_at" field.
func (lc *LikeCreate) SetLikedAt(t time.Time) *LikeCreate {
	lc.mutation.SetLikedAt(t)
	return lc
}

// SetNillableLikedAt sets the "liked_at" field if the given value is not nil.
func (lc *LikeCreate) SetNillableLikedAt(t *time.Time) *LikeCreate {
	if t != nil {
		lc.SetLikedAt(*t)
	}
	return lc
}

// SetUserID sets the "user_id" field.
func (lc *LikeCreate) SetUserID(i int) *LikeCreate {
	lc.mutation.SetUserID(i)
	return lc
}

// SetTweetID sets the "tweet_id" field.
func (lc *LikeCreate) SetTweetID(i int) *LikeCreate {
	lc.mutation.SetTweetID(i)
	return lc
}

// SetUser sets the "user" edge to the User entity.
func (lc *LikeCreate) SetUser(u *User) *LikeCreate {
	return lc.SetUserID(u.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (lc *LikeCreate) SetTweet(t *Tweet) *LikeCreate {
	return lc.SetTweetID(t.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (lc *LikeCreate) Mutation() *LikeMutation {
	return lc.mutation
}

// Save creates the Like in the database.
func (lc *LikeCreate) Save(ctx context.Context) (*Like, error) {
	if err := lc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LikeCreate) SaveX(ctx context.Context) *Like {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LikeCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LikeCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LikeCreate) defaults() error {
	if _, ok := lc.mutation.CreateTime(); !ok {
		if like.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized like.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := like.DefaultCreateTime()
		lc.mutation.SetCreateTime(v)
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		if like.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized like.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := like.DefaultUpdateTime()
		lc.mutation.SetUpdateTime(v)
	}
	if _, ok := lc.mutation.LikedAt(); !ok {
		if like.DefaultLikedAt == nil {
			return fmt.Errorf("ent: uninitialized like.DefaultLikedAt (forgotten import ent/runtime?)")
		}
		v := like.DefaultLikedAt()
		lc.mutation.SetLikedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (lc *LikeCreate) check() error {
	if _, ok := lc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Like.create_time"`)}
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Like.update_time"`)}
	}
	if _, ok := lc.mutation.LikedAt(); !ok {
		return &ValidationError{Name: "liked_at", err: errors.New(`ent: missing required field "Like.liked_at"`)}
	}
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Like.user_id"`)}
	}
	if _, ok := lc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "Like.tweet_id"`)}
	}
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Like.user"`)}
	}
	if _, ok := lc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "Like.tweet"`)}
	}
	return nil
}

func (lc *LikeCreate) sqlSave(ctx context.Context) (*Like, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (lc *LikeCreate) createSpec() (*Like, *sqlgraph.CreateSpec) {
	var (
		_node = &Like{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(like.Table, nil)
	)
	_spec.OnConflict = lc.conflict
	if value, ok := lc.mutation.CreateTime(); ok {
		_spec.SetField(like.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := lc.mutation.UpdateTime(); ok {
		_spec.SetField(like.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := lc.mutation.LikedAt(); ok {
		_spec.SetField(like.FieldLikedAt, field.TypeTime, value)
		_node.LikedAt = value
	}
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   like.TweetTable,
			Columns: []string{like.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TweetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Like.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LikeUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (lc *LikeCreate) OnConflict(opts ...sql.ConflictOption) *LikeUpsertOne {
	lc.conflict = opts
	return &LikeUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lc *LikeCreate) OnConflictColumns(columns ...string) *LikeUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LikeUpsertOne{
		create: lc,
	}
}

type (
	// LikeUpsertOne is the builder for "upsert"-ing
	//  one Like node.
	LikeUpsertOne struct {
		create *LikeCreate
	}

	// LikeUpsert is the "OnConflict" setter.
	LikeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *LikeUpsert) SetUpdateTime(v time.Time) *LikeUpsert {
	u.Set(like.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LikeUpsert) UpdateUpdateTime() *LikeUpsert {
	u.SetExcluded(like.FieldUpdateTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LikeUpsertOne) UpdateNewValues() *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(like.FieldCreateTime)
		}
		if _, exists := u.create.mutation.LikedAt(); exists {
			s.SetIgnore(like.FieldLikedAt)
		}
		if _, exists := u.create.mutation.UserID(); exists {
			s.SetIgnore(like.FieldUserID)
		}
		if _, exists := u.create.mutation.TweetID(); exists {
			s.SetIgnore(like.FieldTweetID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Like.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LikeUpsertOne) Ignore() *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LikeUpsertOne) DoNothing() *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LikeCreate.OnConflict
// documentation for more info.
func (u *LikeUpsertOne) Update(set func(*LikeUpsert)) *LikeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LikeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *LikeUpsertOne) SetUpdateTime(v time.Time) *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LikeUpsertOne) UpdateUpdateTime() *LikeUpsertOne {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *LikeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LikeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LikeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// LikeCreateBulk is the builder for creating many Like entities in bulk.
type LikeCreateBulk struct {
	config
	err      error
	builders []*LikeCreate
	conflict []sql.ConflictOption
}

// Save creates the Like entities in the database.
func (lcb *LikeCreateBulk) Save(ctx context.Context) ([]*Like, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Like, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LikeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LikeCreateBulk) SaveX(ctx context.Context) []*Like {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LikeCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LikeCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Like.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LikeUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (lcb *LikeCreateBulk) OnConflict(opts ...sql.ConflictOption) *LikeUpsertBulk {
	lcb.conflict = opts
	return &LikeUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lcb *LikeCreateBulk) OnConflictColumns(columns ...string) *LikeUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LikeUpsertBulk{
		create: lcb,
	}
}

// LikeUpsertBulk is the builder for "upsert"-ing
// a bulk of Like nodes.
type LikeUpsertBulk struct {
	create *LikeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LikeUpsertBulk) UpdateNewValues() *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(like.FieldCreateTime)
			}
			if _, exists := b.mutation.LikedAt(); exists {
				s.SetIgnore(like.FieldLikedAt)
			}
			if _, exists := b.mutation.UserID(); exists {
				s.SetIgnore(like.FieldUserID)
			}
			if _, exists := b.mutation.TweetID(); exists {
				s.SetIgnore(like.FieldTweetID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Like.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LikeUpsertBulk) Ignore() *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LikeUpsertBulk) DoNothing() *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LikeCreateBulk.OnConflict
// documentation for more info.
func (u *LikeUpsertBulk) Update(set func(*LikeUpsert)) *LikeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LikeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *LikeUpsertBulk) SetUpdateTime(v time.Time) *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LikeUpsertBulk) UpdateUpdateTime() *LikeUpsertBulk {
	return u.Update(func(s *LikeUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *LikeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LikeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LikeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LikeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
