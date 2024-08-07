// Code generated by ent, DO NOT EDIT.

package entity

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lostsnow/keqing/pkg/entity/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// IsBot holds the value of the "is_bot" field.
	IsBot bool `json:"is_bot,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))

	for i := range columns {
		switch columns[i] {
		case user.FieldIsBot:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldUserID:
			values[i] = new(sql.NullInt64)
		case user.FieldUserName, user.FieldFirstName, user.FieldLastName:
			values[i] = new(sql.NullString)
		case user.FieldCreateAt, user.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}

	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}

	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}

			u.ID = int64(value.Int64)
		case user.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				u.UserID = value.Int64
			}
		case user.FieldIsBot:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_bot", values[i])
			} else if value.Valid {
				u.IsBot = value.Bool
			}
		case user.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				u.UserName = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				u.CreateAt = value.Time
			}
		case user.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				u.UpdateAt = value.Time
			}
		}
	}

	return nil
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("entity: User is not a transactional entity")
	}

	u.config.driver = _tx.drv

	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder

	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", u.UserID))
	builder.WriteString(", ")
	builder.WriteString("is_bot=")
	builder.WriteString(fmt.Sprintf("%v", u.IsBot))
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(u.UserName)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(u.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(u.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')

	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
