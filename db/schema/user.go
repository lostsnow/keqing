package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.Bool("is_bot").Default(false),
		field.String("user_name").Default(""),
		// required
		field.String("first_name").Default(""),
		field.String("last_name").Default(""),
		field.Time("create_at").Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("update_at").Default(time.Now).UpdateDefault(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").Unique(),
		index.Fields("user_name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
