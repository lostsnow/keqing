package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ChatOption holds the schema definition for the ChatOption entity.
type ChatOption struct {
	ent.Schema
}

// Fields of the ChatOption.
func (ChatOption) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("chat_id"),
		field.String("key").MaxLen(50),
		field.String("value").Default(""),
		field.Time("create_at").Default(time.Now).Immutable().SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (ChatOption) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chat_id", "key").Unique(),
	}
}

// Edges of the ChatOption.
func (ChatOption) Edges() []ent.Edge {
	return nil
}
