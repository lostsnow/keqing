package db

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ../pkg/entity --feature sql/versioned-migration,sql/upsert,sql/lock,sql/modifier,sql/execquery ./schema
