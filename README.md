# Keqing

Genshin telegram bot

## Database

### Schema

new schema

```shell
go run -mod=mod entgo.io/ent/cmd/ent new --target db/schema <Model>
```

generate entity

```shell
go generate ./db
```

### Migration

install atlas

```shell
go install ariga.io/atlas/cmd/atlas@master
```

update migrations hash

```shell
go run main.go migrate hash
```

generate migrations

```shell
go run main.go migrate generate --name <migration_name>
```

lint migrations

```shell
go run main.go migrate lint --latest 1
```

apply migrations

```shell
go run main.go migrate apply
```

## Translation

install gotext cmd

```shell
go install golang.org/x/text/cmd/gotext@master
```

generate catalog

```shell
go generate -v ./pkg/i18n
```
