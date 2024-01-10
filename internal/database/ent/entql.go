// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/like"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/predicate"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/tweet"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   like.Table,
			Columns: like.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: like.FieldUserID,
				},
				{
					Type:   field.TypeInt,
					Column: like.FieldTweetID,
				},
			},
		},
		Type: "Like",
		Fields: map[string]*sqlgraph.FieldSpec{
			like.FieldCreateTime: {Type: field.TypeTime, Column: like.FieldCreateTime},
			like.FieldUpdateTime: {Type: field.TypeTime, Column: like.FieldUpdateTime},
			like.FieldLikedAt:    {Type: field.TypeTime, Column: like.FieldLikedAt},
			like.FieldUserID:     {Type: field.TypeInt, Column: like.FieldUserID},
			like.FieldTweetID:    {Type: field.TypeInt, Column: like.FieldTweetID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		},
		Type: "Tweet",
		Fields: map[string]*sqlgraph.FieldSpec{
			tweet.FieldText: {Type: field.TypeString, Column: tweet.FieldText},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldAge:  {Type: field.TypeInt, Column: user.FieldAge},
			user.FieldName: {Type: field.TypeString, Column: user.FieldName},
		},
	}
	graph.MustAddE(
		"user",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
		},
		"Like",
		"User",
	)
	graph.MustAddE(
		"tweet",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   like.TweetTable,
			Columns: []string{like.TweetColumn},
			Bidi:    false,
		},
		"Like",
		"Tweet",
	)
	graph.MustAddE(
		"liked_users",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
		},
		"Tweet",
		"User",
	)
	graph.MustAddE(
		"likes",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.LikesTable,
			Columns: []string{tweet.LikesColumn},
			Bidi:    false,
		},
		"Tweet",
		"Like",
	)
	graph.MustAddE(
		"liked_tweets",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.LikedTweetsTable,
			Columns: user.LikedTweetsPrimaryKey,
			Bidi:    false,
		},
		"User",
		"Tweet",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (lq *LikeQuery) addPredicate(pred func(s *sql.Selector)) {
	lq.predicates = append(lq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the LikeQuery builder.
func (lq *LikeQuery) Filter() *LikeFilter {
	return &LikeFilter{config: lq.config, predicateAdder: lq}
}

// addPredicate implements the predicateAdder interface.
func (m *LikeMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the LikeMutation builder.
func (m *LikeMutation) Filter() *LikeFilter {
	return &LikeFilter{config: m.config, predicateAdder: m}
}

// LikeFilter provides a generic filtering capability at runtime for LikeQuery.
type LikeFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *LikeFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *LikeFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(like.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *LikeFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(like.FieldUpdateTime))
}

// WhereLikedAt applies the entql time.Time predicate on the liked_at field.
func (f *LikeFilter) WhereLikedAt(p entql.TimeP) {
	f.Where(p.Field(like.FieldLikedAt))
}

// WhereUserID applies the entql int predicate on the user_id field.
func (f *LikeFilter) WhereUserID(p entql.IntP) {
	f.Where(p.Field(like.FieldUserID))
}

// WhereTweetID applies the entql int predicate on the tweet_id field.
func (f *LikeFilter) WhereTweetID(p entql.IntP) {
	f.Where(p.Field(like.FieldTweetID))
}

// WhereHasUser applies a predicate to check if query has an edge user.
func (f *LikeFilter) WhereHasUser() {
	f.Where(entql.HasEdge("user"))
}

// WhereHasUserWith applies a predicate to check if query has an edge user with a given conditions (other predicates).
func (f *LikeFilter) WhereHasUserWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("user", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasTweet applies a predicate to check if query has an edge tweet.
func (f *LikeFilter) WhereHasTweet() {
	f.Where(entql.HasEdge("tweet"))
}

// WhereHasTweetWith applies a predicate to check if query has an edge tweet with a given conditions (other predicates).
func (f *LikeFilter) WhereHasTweetWith(preds ...predicate.Tweet) {
	f.Where(entql.HasEdgeWith("tweet", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tq *TweetQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TweetQuery builder.
func (tq *TweetQuery) Filter() *TweetFilter {
	return &TweetFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TweetMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TweetMutation builder.
func (m *TweetMutation) Filter() *TweetFilter {
	return &TweetFilter{config: m.config, predicateAdder: m}
}

// TweetFilter provides a generic filtering capability at runtime for TweetQuery.
type TweetFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TweetFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TweetFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(tweet.FieldID))
}

// WhereText applies the entql string predicate on the text field.
func (f *TweetFilter) WhereText(p entql.StringP) {
	f.Where(p.Field(tweet.FieldText))
}

// WhereHasLikedUsers applies a predicate to check if query has an edge liked_users.
func (f *TweetFilter) WhereHasLikedUsers() {
	f.Where(entql.HasEdge("liked_users"))
}

// WhereHasLikedUsersWith applies a predicate to check if query has an edge liked_users with a given conditions (other predicates).
func (f *TweetFilter) WhereHasLikedUsersWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("liked_users", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasLikes applies a predicate to check if query has an edge likes.
func (f *TweetFilter) WhereHasLikes() {
	f.Where(entql.HasEdge("likes"))
}

// WhereHasLikesWith applies a predicate to check if query has an edge likes with a given conditions (other predicates).
func (f *TweetFilter) WhereHasLikesWith(preds ...predicate.Like) {
	f.Where(entql.HasEdgeWith("likes", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *UserFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(user.FieldID))
}

// WhereAge applies the entql int predicate on the age field.
func (f *UserFilter) WhereAge(p entql.IntP) {
	f.Where(p.Field(user.FieldAge))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereHasLikedTweets applies a predicate to check if query has an edge liked_tweets.
func (f *UserFilter) WhereHasLikedTweets() {
	f.Where(entql.HasEdge("liked_tweets"))
}

// WhereHasLikedTweetsWith applies a predicate to check if query has an edge liked_tweets with a given conditions (other predicates).
func (f *UserFilter) WhereHasLikedTweetsWith(preds ...predicate.Tweet) {
	f.Where(entql.HasEdgeWith("liked_tweets", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}