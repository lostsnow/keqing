// Code generated by ent, DO NOT EDIT.

package entity

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lostsnow/keqing/pkg/entity/chatoption"
)

// ChatOption is the model entity for the ChatOption schema.
type ChatOption struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ChatID holds the value of the "chat_id" field.
	ChatID int64 `json:"chat_id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChatOption) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))

	for i := range columns {
		switch columns[i] {
		case chatoption.FieldID, chatoption.FieldChatID:
			values[i] = new(sql.NullInt64)
		case chatoption.FieldKey, chatoption.FieldValue:
			values[i] = new(sql.NullString)
		case chatoption.FieldCreateAt, chatoption.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChatOption", columns[i])
		}
	}

	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChatOption fields.
func (co *ChatOption) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}

	for i := range columns {
		switch columns[i] {
		case chatoption.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}

			co.ID = int64(value.Int64)
		case chatoption.FieldChatID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chat_id", values[i])
			} else if value.Valid {
				co.ChatID = value.Int64
			}
		case chatoption.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				co.Key = value.String
			}
		case chatoption.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				co.Value = value.String
			}
		case chatoption.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				co.CreateAt = value.Time
			}
		case chatoption.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				co.UpdateAt = value.Time
			}
		}
	}

	return nil
}

// Update returns a builder for updating this ChatOption.
// Note that you need to call ChatOption.Unwrap() before calling this method if this ChatOption
// was returned from a transaction, and the transaction was committed or rolled back.
func (co *ChatOption) Update() *ChatOptionUpdateOne {
	return NewChatOptionClient(co.config).UpdateOne(co)
}

// Unwrap unwraps the ChatOption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (co *ChatOption) Unwrap() *ChatOption {
	_tx, ok := co.config.driver.(*txDriver)
	if !ok {
		panic("entity: ChatOption is not a transactional entity")
	}

	co.config.driver = _tx.drv

	return co
}

// String implements the fmt.Stringer.
func (co *ChatOption) String() string {
	var builder strings.Builder

	builder.WriteString("ChatOption(")
	builder.WriteString(fmt.Sprintf("id=%v, ", co.ID))
	builder.WriteString("chat_id=")
	builder.WriteString(fmt.Sprintf("%v", co.ChatID))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(co.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(co.Value)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(co.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(co.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')

	return builder.String()
}

// ChatOptions is a parsable slice of ChatOption.
type ChatOptions []*ChatOption
