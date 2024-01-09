// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/like"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/tweet"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/user"
)

// Like is the model entity for the Like schema.
type Like struct {
	config `json:"-"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// test
	LikedAt time.Time `json:"liked_at,omitempty"`
	// test
	UserID int `json:"user_id,omitempty"`
	// test
	TweetID int `json:"tweet_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LikeQuery when eager-loading is set.
	Edges        LikeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LikeEdges holds the relations/edges for other nodes in the graph.
type LikeEdges struct {
	// test
	User *User `json:"user,omitempty"`
	// test
	Tweet *Tweet `json:"tweet,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LikeEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TweetOrErr returns the Tweet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LikeEdges) TweetOrErr() (*Tweet, error) {
	if e.loadedTypes[1] {
		if e.Tweet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tweet.Label}
		}
		return e.Tweet, nil
	}
	return nil, &NotLoadedError{edge: "tweet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Like) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case like.FieldUserID, like.FieldTweetID:
			values[i] = new(sql.NullInt64)
		case like.FieldCreateTime, like.FieldUpdateTime, like.FieldLikedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Like fields.
func (l *Like) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case like.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				l.CreateTime = value.Time
			}
		case like.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				l.UpdateTime = value.Time
			}
		case like.FieldLikedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field liked_at", values[i])
			} else if value.Valid {
				l.LikedAt = value.Time
			}
		case like.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				l.UserID = int(value.Int64)
			}
		case like.FieldTweetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_id", values[i])
			} else if value.Valid {
				l.TweetID = int(value.Int64)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Like.
// This includes values selected through modifiers, order, etc.
func (l *Like) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Like entity.
func (l *Like) QueryUser() *UserQuery {
	return NewLikeClient(l.config).QueryUser(l)
}

// QueryTweet queries the "tweet" edge of the Like entity.
func (l *Like) QueryTweet() *TweetQuery {
	return NewLikeClient(l.config).QueryTweet(l)
}

// Update returns a builder for updating this Like.
// Note that you need to call Like.Unwrap() before calling this method if this Like
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Like) Update() *LikeUpdateOne {
	return NewLikeClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Like entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Like) Unwrap() *Like {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Like is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Like) String() string {
	var builder strings.Builder
	builder.WriteString("Like(")
	builder.WriteString("create_time=")
	builder.WriteString(l.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(l.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("liked_at=")
	builder.WriteString(l.LikedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", l.UserID))
	builder.WriteString(", ")
	builder.WriteString("tweet_id=")
	builder.WriteString(fmt.Sprintf("%v", l.TweetID))
	builder.WriteByte(')')
	return builder.String()
}

// Likes is a parsable slice of Like.
type Likes []*Like
