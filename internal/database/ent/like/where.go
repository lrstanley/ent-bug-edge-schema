// Code generated by ent, DO NOT EDIT.

package like

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/predicate"
)

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUpdateTime, v))
}

// LikedAt applies equality check predicate on the "liked_at" field. It's identical to LikedAtEQ.
func LikedAt(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldLikedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUserID, v))
}

// TweetID applies equality check predicate on the "tweet_id" field. It's identical to TweetIDEQ.
func TweetID(v int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldTweetID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldUpdateTime, v))
}

// LikedAtEQ applies the EQ predicate on the "liked_at" field.
func LikedAtEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldLikedAt, v))
}

// LikedAtNEQ applies the NEQ predicate on the "liked_at" field.
func LikedAtNEQ(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldLikedAt, v))
}

// LikedAtIn applies the In predicate on the "liked_at" field.
func LikedAtIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldLikedAt, vs...))
}

// LikedAtNotIn applies the NotIn predicate on the "liked_at" field.
func LikedAtNotIn(vs ...time.Time) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldLikedAt, vs...))
}

// LikedAtGT applies the GT predicate on the "liked_at" field.
func LikedAtGT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGT(FieldLikedAt, v))
}

// LikedAtGTE applies the GTE predicate on the "liked_at" field.
func LikedAtGTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldGTE(FieldLikedAt, v))
}

// LikedAtLT applies the LT predicate on the "liked_at" field.
func LikedAtLT(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLT(FieldLikedAt, v))
}

// LikedAtLTE applies the LTE predicate on the "liked_at" field.
func LikedAtLTE(v time.Time) predicate.Like {
	return predicate.Like(sql.FieldLTE(FieldLikedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldUserID, vs...))
}

// TweetIDEQ applies the EQ predicate on the "tweet_id" field.
func TweetIDEQ(v int) predicate.Like {
	return predicate.Like(sql.FieldEQ(FieldTweetID, v))
}

// TweetIDNEQ applies the NEQ predicate on the "tweet_id" field.
func TweetIDNEQ(v int) predicate.Like {
	return predicate.Like(sql.FieldNEQ(FieldTweetID, v))
}

// TweetIDIn applies the In predicate on the "tweet_id" field.
func TweetIDIn(vs ...int) predicate.Like {
	return predicate.Like(sql.FieldIn(FieldTweetID, vs...))
}

// TweetIDNotIn applies the NotIn predicate on the "tweet_id" field.
func TweetIDNotIn(vs ...int) predicate.Like {
	return predicate.Like(sql.FieldNotIn(FieldTweetID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTweet applies the HasEdge predicate on the "tweet" edge.
func HasTweet() predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, TweetColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTweetWith applies the HasEdge predicate on the "tweet" edge with a given conditions (other predicates).
func HasTweetWith(preds ...predicate.Tweet) predicate.Like {
	return predicate.Like(func(s *sql.Selector) {
		step := newTweetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Like) predicate.Like {
	return predicate.Like(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Like) predicate.Like {
	return predicate.Like(sql.NotPredicates(p))
}