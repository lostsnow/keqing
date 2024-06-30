// Code generated by ent, DO NOT EDIT.

package gameaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lostsnow/keqing/pkg/entity/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldUserID, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldAccountID, v))
}

// GameToken applies equality check predicate on the "game_token" field. It's identical to GameTokenEQ.
func GameToken(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldGameToken, v))
}

// CookieToken applies equality check predicate on the "cookie_token" field. It's identical to CookieTokenEQ.
func CookieToken(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldCookieToken, v))
}

// Stoken applies equality check predicate on the "stoken" field. It's identical to StokenEQ.
func Stoken(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldStoken, v))
}

// Mid applies equality check predicate on the "mid" field. It's identical to MidEQ.
func Mid(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldMid, v))
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldUpdateAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldUserID, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldAccountID, v))
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContains(FieldAccountID, v))
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasPrefix(FieldAccountID, v))
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasSuffix(FieldAccountID, v))
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEqualFold(FieldAccountID, v))
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContainsFold(FieldAccountID, v))
}

// GameTokenEQ applies the EQ predicate on the "game_token" field.
func GameTokenEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldGameToken, v))
}

// GameTokenNEQ applies the NEQ predicate on the "game_token" field.
func GameTokenNEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldGameToken, v))
}

// GameTokenIn applies the In predicate on the "game_token" field.
func GameTokenIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldGameToken, vs...))
}

// GameTokenNotIn applies the NotIn predicate on the "game_token" field.
func GameTokenNotIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldGameToken, vs...))
}

// GameTokenGT applies the GT predicate on the "game_token" field.
func GameTokenGT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldGameToken, v))
}

// GameTokenGTE applies the GTE predicate on the "game_token" field.
func GameTokenGTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldGameToken, v))
}

// GameTokenLT applies the LT predicate on the "game_token" field.
func GameTokenLT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldGameToken, v))
}

// GameTokenLTE applies the LTE predicate on the "game_token" field.
func GameTokenLTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldGameToken, v))
}

// GameTokenContains applies the Contains predicate on the "game_token" field.
func GameTokenContains(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContains(FieldGameToken, v))
}

// GameTokenHasPrefix applies the HasPrefix predicate on the "game_token" field.
func GameTokenHasPrefix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasPrefix(FieldGameToken, v))
}

// GameTokenHasSuffix applies the HasSuffix predicate on the "game_token" field.
func GameTokenHasSuffix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasSuffix(FieldGameToken, v))
}

// GameTokenEqualFold applies the EqualFold predicate on the "game_token" field.
func GameTokenEqualFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEqualFold(FieldGameToken, v))
}

// GameTokenContainsFold applies the ContainsFold predicate on the "game_token" field.
func GameTokenContainsFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContainsFold(FieldGameToken, v))
}

// CookieTokenEQ applies the EQ predicate on the "cookie_token" field.
func CookieTokenEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldCookieToken, v))
}

// CookieTokenNEQ applies the NEQ predicate on the "cookie_token" field.
func CookieTokenNEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldCookieToken, v))
}

// CookieTokenIn applies the In predicate on the "cookie_token" field.
func CookieTokenIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldCookieToken, vs...))
}

// CookieTokenNotIn applies the NotIn predicate on the "cookie_token" field.
func CookieTokenNotIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldCookieToken, vs...))
}

// CookieTokenGT applies the GT predicate on the "cookie_token" field.
func CookieTokenGT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldCookieToken, v))
}

// CookieTokenGTE applies the GTE predicate on the "cookie_token" field.
func CookieTokenGTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldCookieToken, v))
}

// CookieTokenLT applies the LT predicate on the "cookie_token" field.
func CookieTokenLT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldCookieToken, v))
}

// CookieTokenLTE applies the LTE predicate on the "cookie_token" field.
func CookieTokenLTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldCookieToken, v))
}

// CookieTokenContains applies the Contains predicate on the "cookie_token" field.
func CookieTokenContains(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContains(FieldCookieToken, v))
}

// CookieTokenHasPrefix applies the HasPrefix predicate on the "cookie_token" field.
func CookieTokenHasPrefix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasPrefix(FieldCookieToken, v))
}

// CookieTokenHasSuffix applies the HasSuffix predicate on the "cookie_token" field.
func CookieTokenHasSuffix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasSuffix(FieldCookieToken, v))
}

// CookieTokenEqualFold applies the EqualFold predicate on the "cookie_token" field.
func CookieTokenEqualFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEqualFold(FieldCookieToken, v))
}

// CookieTokenContainsFold applies the ContainsFold predicate on the "cookie_token" field.
func CookieTokenContainsFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContainsFold(FieldCookieToken, v))
}

// StokenEQ applies the EQ predicate on the "stoken" field.
func StokenEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldStoken, v))
}

// StokenNEQ applies the NEQ predicate on the "stoken" field.
func StokenNEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldStoken, v))
}

// StokenIn applies the In predicate on the "stoken" field.
func StokenIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldStoken, vs...))
}

// StokenNotIn applies the NotIn predicate on the "stoken" field.
func StokenNotIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldStoken, vs...))
}

// StokenGT applies the GT predicate on the "stoken" field.
func StokenGT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldStoken, v))
}

// StokenGTE applies the GTE predicate on the "stoken" field.
func StokenGTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldStoken, v))
}

// StokenLT applies the LT predicate on the "stoken" field.
func StokenLT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldStoken, v))
}

// StokenLTE applies the LTE predicate on the "stoken" field.
func StokenLTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldStoken, v))
}

// StokenContains applies the Contains predicate on the "stoken" field.
func StokenContains(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContains(FieldStoken, v))
}

// StokenHasPrefix applies the HasPrefix predicate on the "stoken" field.
func StokenHasPrefix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasPrefix(FieldStoken, v))
}

// StokenHasSuffix applies the HasSuffix predicate on the "stoken" field.
func StokenHasSuffix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasSuffix(FieldStoken, v))
}

// StokenEqualFold applies the EqualFold predicate on the "stoken" field.
func StokenEqualFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEqualFold(FieldStoken, v))
}

// StokenContainsFold applies the ContainsFold predicate on the "stoken" field.
func StokenContainsFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContainsFold(FieldStoken, v))
}

// MidEQ applies the EQ predicate on the "mid" field.
func MidEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldMid, v))
}

// MidNEQ applies the NEQ predicate on the "mid" field.
func MidNEQ(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldMid, v))
}

// MidIn applies the In predicate on the "mid" field.
func MidIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldMid, vs...))
}

// MidNotIn applies the NotIn predicate on the "mid" field.
func MidNotIn(vs ...string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldMid, vs...))
}

// MidGT applies the GT predicate on the "mid" field.
func MidGT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldMid, v))
}

// MidGTE applies the GTE predicate on the "mid" field.
func MidGTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldMid, v))
}

// MidLT applies the LT predicate on the "mid" field.
func MidLT(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldMid, v))
}

// MidLTE applies the LTE predicate on the "mid" field.
func MidLTE(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldMid, v))
}

// MidContains applies the Contains predicate on the "mid" field.
func MidContains(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContains(FieldMid, v))
}

// MidHasPrefix applies the HasPrefix predicate on the "mid" field.
func MidHasPrefix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasPrefix(FieldMid, v))
}

// MidHasSuffix applies the HasSuffix predicate on the "mid" field.
func MidHasSuffix(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldHasSuffix(FieldMid, v))
}

// MidEqualFold applies the EqualFold predicate on the "mid" field.
func MidEqualFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEqualFold(FieldMid, v))
}

// MidContainsFold applies the ContainsFold predicate on the "mid" field.
func MidContainsFold(v string) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldContainsFold(FieldMid, v))
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.GameAccount {
	return predicate.GameAccount(sql.FieldLTE(FieldUpdateAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GameAccount) predicate.GameAccount {
	return predicate.GameAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}

		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GameAccount) predicate.GameAccount {
	return predicate.GameAccount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)

		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}

			p(s1)
		}

		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GameAccount) predicate.GameAccount {
	return predicate.GameAccount(func(s *sql.Selector) {
		p(s.Not())
	})
}
