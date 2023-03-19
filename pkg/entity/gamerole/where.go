// Code generated by ent, DO NOT EDIT.

package gamerole

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lostsnow/keqing/pkg/entity/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldUserID, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldAccountID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldRoleID, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldLevel, v))
}

// Region applies equality check predicate on the "region" field. It's identical to RegionEQ.
func Region(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldRegion, v))
}

// RegionName applies equality check predicate on the "region_name" field. It's identical to RegionNameEQ.
func RegionName(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldRegionName, v))
}

// NickName applies equality check predicate on the "nick_name" field. It's identical to NickNameEQ.
func NickName(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldNickName, v))
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldUpdateAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldUserID, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldAccountID, v))
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContains(FieldAccountID, v))
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasPrefix(FieldAccountID, v))
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasSuffix(FieldAccountID, v))
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEqualFold(FieldAccountID, v))
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContainsFold(FieldAccountID, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContains(FieldRoleID, v))
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasPrefix(FieldRoleID, v))
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasSuffix(FieldRoleID, v))
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEqualFold(FieldRoleID, v))
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContainsFold(FieldRoleID, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldLevel, v))
}

// RegionEQ applies the EQ predicate on the "region" field.
func RegionEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldRegion, v))
}

// RegionNEQ applies the NEQ predicate on the "region" field.
func RegionNEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldRegion, v))
}

// RegionIn applies the In predicate on the "region" field.
func RegionIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldRegion, vs...))
}

// RegionNotIn applies the NotIn predicate on the "region" field.
func RegionNotIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldRegion, vs...))
}

// RegionGT applies the GT predicate on the "region" field.
func RegionGT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldRegion, v))
}

// RegionGTE applies the GTE predicate on the "region" field.
func RegionGTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldRegion, v))
}

// RegionLT applies the LT predicate on the "region" field.
func RegionLT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldRegion, v))
}

// RegionLTE applies the LTE predicate on the "region" field.
func RegionLTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldRegion, v))
}

// RegionContains applies the Contains predicate on the "region" field.
func RegionContains(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContains(FieldRegion, v))
}

// RegionHasPrefix applies the HasPrefix predicate on the "region" field.
func RegionHasPrefix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasPrefix(FieldRegion, v))
}

// RegionHasSuffix applies the HasSuffix predicate on the "region" field.
func RegionHasSuffix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasSuffix(FieldRegion, v))
}

// RegionEqualFold applies the EqualFold predicate on the "region" field.
func RegionEqualFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEqualFold(FieldRegion, v))
}

// RegionContainsFold applies the ContainsFold predicate on the "region" field.
func RegionContainsFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContainsFold(FieldRegion, v))
}

// RegionNameEQ applies the EQ predicate on the "region_name" field.
func RegionNameEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldRegionName, v))
}

// RegionNameNEQ applies the NEQ predicate on the "region_name" field.
func RegionNameNEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldRegionName, v))
}

// RegionNameIn applies the In predicate on the "region_name" field.
func RegionNameIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldRegionName, vs...))
}

// RegionNameNotIn applies the NotIn predicate on the "region_name" field.
func RegionNameNotIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldRegionName, vs...))
}

// RegionNameGT applies the GT predicate on the "region_name" field.
func RegionNameGT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldRegionName, v))
}

// RegionNameGTE applies the GTE predicate on the "region_name" field.
func RegionNameGTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldRegionName, v))
}

// RegionNameLT applies the LT predicate on the "region_name" field.
func RegionNameLT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldRegionName, v))
}

// RegionNameLTE applies the LTE predicate on the "region_name" field.
func RegionNameLTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldRegionName, v))
}

// RegionNameContains applies the Contains predicate on the "region_name" field.
func RegionNameContains(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContains(FieldRegionName, v))
}

// RegionNameHasPrefix applies the HasPrefix predicate on the "region_name" field.
func RegionNameHasPrefix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasPrefix(FieldRegionName, v))
}

// RegionNameHasSuffix applies the HasSuffix predicate on the "region_name" field.
func RegionNameHasSuffix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasSuffix(FieldRegionName, v))
}

// RegionNameEqualFold applies the EqualFold predicate on the "region_name" field.
func RegionNameEqualFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEqualFold(FieldRegionName, v))
}

// RegionNameContainsFold applies the ContainsFold predicate on the "region_name" field.
func RegionNameContainsFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContainsFold(FieldRegionName, v))
}

// NickNameEQ applies the EQ predicate on the "nick_name" field.
func NickNameEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldNickName, v))
}

// NickNameNEQ applies the NEQ predicate on the "nick_name" field.
func NickNameNEQ(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldNickName, v))
}

// NickNameIn applies the In predicate on the "nick_name" field.
func NickNameIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldNickName, vs...))
}

// NickNameNotIn applies the NotIn predicate on the "nick_name" field.
func NickNameNotIn(vs ...string) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldNickName, vs...))
}

// NickNameGT applies the GT predicate on the "nick_name" field.
func NickNameGT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldNickName, v))
}

// NickNameGTE applies the GTE predicate on the "nick_name" field.
func NickNameGTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldNickName, v))
}

// NickNameLT applies the LT predicate on the "nick_name" field.
func NickNameLT(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldNickName, v))
}

// NickNameLTE applies the LTE predicate on the "nick_name" field.
func NickNameLTE(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldNickName, v))
}

// NickNameContains applies the Contains predicate on the "nick_name" field.
func NickNameContains(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContains(FieldNickName, v))
}

// NickNameHasPrefix applies the HasPrefix predicate on the "nick_name" field.
func NickNameHasPrefix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasPrefix(FieldNickName, v))
}

// NickNameHasSuffix applies the HasSuffix predicate on the "nick_name" field.
func NickNameHasSuffix(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldHasSuffix(FieldNickName, v))
}

// NickNameEqualFold applies the EqualFold predicate on the "nick_name" field.
func NickNameEqualFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldEqualFold(FieldNickName, v))
}

// NickNameContainsFold applies the ContainsFold predicate on the "nick_name" field.
func NickNameContainsFold(v string) predicate.GameRole {
	return predicate.GameRole(sql.FieldContainsFold(FieldNickName, v))
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.GameRole {
	return predicate.GameRole(sql.FieldLTE(FieldUpdateAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GameRole) predicate.GameRole {
	return predicate.GameRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GameRole) predicate.GameRole {
	return predicate.GameRole(func(s *sql.Selector) {
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
func Not(p predicate.GameRole) predicate.GameRole {
	return predicate.GameRole(func(s *sql.Selector) {
		p(s.Not())
	})
}