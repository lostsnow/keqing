package db

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ../pkg/entity --feature sql/versioned-migration ./schema
