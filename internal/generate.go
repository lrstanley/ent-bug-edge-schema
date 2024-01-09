package internal

//go:generate sh -c "mkdir -p ./graphql/schema ./database/schema"
//go:generate go run -mod=mod database/generate.go
//go:generate go run -mod=mod github.com/99designs/gqlgen generate --config graphql/gqlgen.yml
