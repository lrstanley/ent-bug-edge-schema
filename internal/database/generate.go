//go:build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	egq, err := entgql.NewExtension(
		entgql.WithConfigPath("./graphql/gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./graphql/schema/generated.gql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		panic(err)
	}

	err = entc.Generate(
		"./database/schema",
		&gen.Config{
			Target:  "./database/ent",
			Schema:  "github.com/lrstanley/ent-bug-edge-schema/internal/database/schema",
			Package: "github.com/lrstanley/ent-bug-edge-schema/internal/database/ent",
			Features: []gen.Feature{
				gen.FeaturePrivacy,
				gen.FeatureEntQL,
				gen.FeatureSnapshot,
				gen.FeatureUpsert,
			},
		},
		entc.Extensions(egq),
	)
	if err != nil {
		panic(err)
	}
}
