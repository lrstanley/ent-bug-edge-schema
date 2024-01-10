// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/tweet"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/user"
)

// TweetCreate is the builder for creating a Tweet entity.
type TweetCreate struct {
	config
	mutation *TweetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetText sets the "text" field.
func (tc *TweetCreate) SetText(s string) *TweetCreate {
	tc.mutation.SetText(s)
	return tc
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (tc *TweetCreate) AddLikedUserIDs(ids ...int) *TweetCreate {
	tc.mutation.AddLikedUserIDs(ids...)
	return tc
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (tc *TweetCreate) AddLikedUsers(u ...*User) *TweetCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddLikedUserIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (tc *TweetCreate) Mutation() *TweetMutation {
	return tc.mutation
}

// Save creates the Tweet in the database.
func (tc *TweetCreate) Save(ctx context.Context) (*Tweet, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TweetCreate) SaveX(ctx context.Context) *Tweet {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TweetCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TweetCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TweetCreate) check() error {
	if _, ok := tc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Tweet.text"`)}
	}
	return nil
}

func (tc *TweetCreate) sqlSave(ctx context.Context) (*Tweet, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TweetCreate) createSpec() (*Tweet, *sqlgraph.CreateSpec) {
	var (
		_node = &Tweet{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tweet.Table, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if nodes := tc.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &LikeCreate{config: tc.config, mutation: newLikeMutation(tc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tweet.Create().
//		SetText(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TweetUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (tc *TweetCreate) OnConflict(opts ...sql.ConflictOption) *TweetUpsertOne {
	tc.conflict = opts
	return &TweetUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tweet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TweetCreate) OnConflictColumns(columns ...string) *TweetUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TweetUpsertOne{
		create: tc,
	}
}

type (
	// TweetUpsertOne is the builder for "upsert"-ing
	//  one Tweet node.
	TweetUpsertOne struct {
		create *TweetCreate
	}

	// TweetUpsert is the "OnConflict" setter.
	TweetUpsert struct {
		*sql.UpdateSet
	}
)

// SetText sets the "text" field.
func (u *TweetUpsert) SetText(v string) *TweetUpsert {
	u.Set(tweet.FieldText, v)
	return u
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *TweetUpsert) UpdateText() *TweetUpsert {
	u.SetExcluded(tweet.FieldText)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Tweet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TweetUpsertOne) UpdateNewValues() *TweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tweet.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TweetUpsertOne) Ignore() *TweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TweetUpsertOne) DoNothing() *TweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TweetCreate.OnConflict
// documentation for more info.
func (u *TweetUpsertOne) Update(set func(*TweetUpsert)) *TweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TweetUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *TweetUpsertOne) SetText(v string) *TweetUpsertOne {
	return u.Update(func(s *TweetUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *TweetUpsertOne) UpdateText() *TweetUpsertOne {
	return u.Update(func(s *TweetUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *TweetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TweetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TweetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TweetUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TweetUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TweetCreateBulk is the builder for creating many Tweet entities in bulk.
type TweetCreateBulk struct {
	config
	err      error
	builders []*TweetCreate
	conflict []sql.ConflictOption
}

// Save creates the Tweet entities in the database.
func (tcb *TweetCreateBulk) Save(ctx context.Context) ([]*Tweet, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tweet, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TweetMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TweetCreateBulk) SaveX(ctx context.Context) []*Tweet {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TweetCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TweetCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tweet.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TweetUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (tcb *TweetCreateBulk) OnConflict(opts ...sql.ConflictOption) *TweetUpsertBulk {
	tcb.conflict = opts
	return &TweetUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tweet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TweetCreateBulk) OnConflictColumns(columns ...string) *TweetUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TweetUpsertBulk{
		create: tcb,
	}
}

// TweetUpsertBulk is the builder for "upsert"-ing
// a bulk of Tweet nodes.
type TweetUpsertBulk struct {
	create *TweetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tweet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TweetUpsertBulk) UpdateNewValues() *TweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tweet.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TweetUpsertBulk) Ignore() *TweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TweetUpsertBulk) DoNothing() *TweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TweetCreateBulk.OnConflict
// documentation for more info.
func (u *TweetUpsertBulk) Update(set func(*TweetUpsert)) *TweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TweetUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *TweetUpsertBulk) SetText(v string) *TweetUpsertBulk {
	return u.Update(func(s *TweetUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *TweetUpsertBulk) UpdateText() *TweetUpsertBulk {
	return u.Update(func(s *TweetUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *TweetUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TweetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TweetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TweetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}