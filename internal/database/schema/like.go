package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/lrstanley/ent-bug-edge-schema/internal/database/ent/privacy"
)

// Like holds the edge schema definition for the Like edge.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Time("liked_at").Default(time.Now).Immutable().Comment("test"),
		field.Int("user_id").Unique().Immutable().Comment("test").Annotations(entgql.OrderField("USER_ID")),
		field.Int("tweet_id").Unique().Immutable().Comment("test").Annotations(entgql.OrderField("TWEET_ID")),
	}
}

func (Like) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
		mixin.UpdateTime{},
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Immutable().Required().Field("user_id").Comment("test"),
		edge.To("tweet", Tweet.Type).Unique().Immutable().Required().Field("tweet_id").Comment("test"),
	}
}

func (Like) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "tweet_id"),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
