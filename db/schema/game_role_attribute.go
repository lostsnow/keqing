package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// GameRoleAttribute holds the schema definition for the GameRoleAttribute entity.
type GameRoleAttribute struct {
	ent.Schema
}

// Fields of the GameRoleAttribute.
func (GameRoleAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("account_id").MaxLen(30),
		field.String("role_id").MaxLen(30),
		field.String("name").MaxLen(100),
		field.Int("type").Default(0),
		field.String("value").Default(""),
		field.Time("create_at").Default(time.Now).Immutable().SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (GameRoleAttribute) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "account_id", "role_id", "name").Unique(),
		index.Fields("account_id"),
		index.Fields("role_id"),
		index.Fields("name"),
	}
}

// Edges of the GameRoleAttribute.
func (GameRoleAttribute) Edges() []ent.Edge {
	return nil
}
