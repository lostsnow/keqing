// Code generated by ent, DO NOT EDIT.

package entity

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lostsnow/keqing/pkg/entity/gameaccount"
)

// GameAccount is the model entity for the GameAccount schema.
type GameAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID string `json:"account_id,omitempty"`
	// GameToken holds the value of the "game_token" field.
	GameToken string `json:"game_token,omitempty"`
	// CookieToken holds the value of the "cookie_token" field.
	CookieToken string `json:"cookie_token,omitempty"`
	// Stoken holds the value of the "stoken" field.
	Stoken string `json:"stoken,omitempty"`
	// Mid holds the value of the "mid" field.
	Mid string `json:"mid,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GameAccount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))

	for i := range columns {
		switch columns[i] {
		case gameaccount.FieldID, gameaccount.FieldUserID:
			values[i] = new(sql.NullInt64)
		case gameaccount.FieldAccountID, gameaccount.FieldGameToken, gameaccount.FieldCookieToken, gameaccount.FieldStoken, gameaccount.FieldMid:
			values[i] = new(sql.NullString)
		case gameaccount.FieldCreateAt, gameaccount.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GameAccount", columns[i])
		}
	}

	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GameAccount fields.
func (ga *GameAccount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}

	for i := range columns {
		switch columns[i] {
		case gameaccount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}

			ga.ID = int64(value.Int64)
		case gameaccount.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ga.UserID = value.Int64
			}
		case gameaccount.FieldAccountID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				ga.AccountID = value.String
			}
		case gameaccount.FieldGameToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field game_token", values[i])
			} else if value.Valid {
				ga.GameToken = value.String
			}
		case gameaccount.FieldCookieToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cookie_token", values[i])
			} else if value.Valid {
				ga.CookieToken = value.String
			}
		case gameaccount.FieldStoken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stoken", values[i])
			} else if value.Valid {
				ga.Stoken = value.String
			}
		case gameaccount.FieldMid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mid", values[i])
			} else if value.Valid {
				ga.Mid = value.String
			}
		case gameaccount.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ga.CreateAt = value.Time
			}
		case gameaccount.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ga.UpdateAt = value.Time
			}
		}
	}

	return nil
}

// Update returns a builder for updating this GameAccount.
// Note that you need to call GameAccount.Unwrap() before calling this method if this GameAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *GameAccount) Update() *GameAccountUpdateOne {
	return NewGameAccountClient(ga.config).UpdateOne(ga)
}

// Unwrap unwraps the GameAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *GameAccount) Unwrap() *GameAccount {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("entity: GameAccount is not a transactional entity")
	}

	ga.config.driver = _tx.drv

	return ga
}

// String implements the fmt.Stringer.
func (ga *GameAccount) String() string {
	var builder strings.Builder

	builder.WriteString("GameAccount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ga.UserID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(ga.AccountID)
	builder.WriteString(", ")
	builder.WriteString("game_token=")
	builder.WriteString(ga.GameToken)
	builder.WriteString(", ")
	builder.WriteString("cookie_token=")
	builder.WriteString(ga.CookieToken)
	builder.WriteString(", ")
	builder.WriteString("stoken=")
	builder.WriteString(ga.Stoken)
	builder.WriteString(", ")
	builder.WriteString("mid=")
	builder.WriteString(ga.Mid)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(ga.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(ga.UpdateAt.Format(time.ANSIC))
	builder.WriteByte(')')

	return builder.String()
}

// GameAccounts is a parsable slice of GameAccount.
type GameAccounts []*GameAccount
