package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// GameAccount holds the schema definition for the GameAccount entity.
type GameAccount struct {
	ent.Schema
}

// Fields of the GameAccount.
func (GameAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("account_id").MaxLen(30),
		field.String("game_token").MaxLen(100).Default(""),
		field.String("cookie_token").MaxLen(100).Default(""),
		field.String("stoken").Default(""),
		field.String("mid").MaxLen(50).Default(""),
		field.Time("create_at").Default(time.Now).Immutable().SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (GameAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "account_id").Unique(),
		index.Fields("account_id"),
	}
}

// Edges of the GameAccount.
func (GameAccount) Edges() []ent.Edge {
	return nil
}
