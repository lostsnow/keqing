// Code generated by ent, DO NOT EDIT.

package chatoption

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lostsnow/keqing/pkg/entity/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLTE(FieldID, id))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldChatID, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldValue, v))
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldUpdateAt, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v int64) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLTE(FieldChatID, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldContainsFold(FieldValue, v))
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.ChatOption {
	return predicate.ChatOption(sql.FieldLTE(FieldUpdateAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChatOption) predicate.ChatOption {
	return predicate.ChatOption(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}

		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChatOption) predicate.ChatOption {
	return predicate.ChatOption(func(s *sql.Selector) {
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
func Not(p predicate.ChatOption) predicate.ChatOption {
	return predicate.ChatOption(func(s *sql.Selector) {
		p(s.Not())
	})
}
