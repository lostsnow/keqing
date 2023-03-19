package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// GameRole holds the schema definition for the GameRole entity.
type GameRole struct {
	ent.Schema
}

// Fields of the GameRole.
func (GameRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("account_id").MaxLen(30),
		field.String("role_id").MaxLen(30),
		field.Int("level").Default(0),
		field.String("region").MaxLen(30).Default(""),
		field.String("region_name").MaxLen(30).Default(""),
		field.String("nick_name").MaxLen(100).Default(""),
		field.Time("create_at").Default(time.Now).Immutable().SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (GameRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "account_id", "role_id").Unique(),
		index.Fields("account_id"),
		index.Fields("role_id"),
		index.Fields("region"),
	}
}

// Edges of the GameRole.
func (GameRole) Edges() []ent.Edge {
	return nil
}
