package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("chat_id"),
		// private/group/supergroup/channel
		field.String("type").Default(""),
		field.Bool("is_forum").Default(false),
		field.String("title").Default(""),
		field.String("user_name").Default(""),
		field.String("first_name").Default(""),
		field.String("last_name").Default(""),
		field.Text("description").Default("").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.Time("create_at").Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (Chat) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chat_id").Unique(),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return nil
}
