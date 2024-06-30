// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lostsnow/keqing/pkg/entity/gameroleattribute"
	"github.com/lostsnow/keqing/pkg/entity/predicate"
)

// GameRoleAttributeUpdate is the builder for updating GameRoleAttribute entities.
type GameRoleAttributeUpdate struct {
	config
	hooks     []Hook
	mutation  *GameRoleAttributeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GameRoleAttributeUpdate builder.
func (grau *GameRoleAttributeUpdate) Where(ps ...predicate.GameRoleAttribute) *GameRoleAttributeUpdate {
	grau.mutation.Where(ps...)
	return grau
}

// SetUserID sets the "user_id" field.
func (grau *GameRoleAttributeUpdate) SetUserID(i int64) *GameRoleAttributeUpdate {
	grau.mutation.ResetUserID()
	grau.mutation.SetUserID(i)

	return grau
}

// AddUserID adds i to the "user_id" field.
func (grau *GameRoleAttributeUpdate) AddUserID(i int64) *GameRoleAttributeUpdate {
	grau.mutation.AddUserID(i)
	return grau
}

// SetAccountID sets the "account_id" field.
func (grau *GameRoleAttributeUpdate) SetAccountID(s string) *GameRoleAttributeUpdate {
	grau.mutation.SetAccountID(s)
	return grau
}

// SetRoleID sets the "role_id" field.
func (grau *GameRoleAttributeUpdate) SetRoleID(s string) *GameRoleAttributeUpdate {
	grau.mutation.SetRoleID(s)
	return grau
}

// SetName sets the "name" field.
func (grau *GameRoleAttributeUpdate) SetName(s string) *GameRoleAttributeUpdate {
	grau.mutation.SetName(s)
	return grau
}

// SetType sets the "type" field.
func (grau *GameRoleAttributeUpdate) SetType(i int) *GameRoleAttributeUpdate {
	grau.mutation.ResetType()
	grau.mutation.SetType(i)

	return grau
}

// SetNillableType sets the "type" field if the given value is not nil.
func (grau *GameRoleAttributeUpdate) SetNillableType(i *int) *GameRoleAttributeUpdate {
	if i != nil {
		grau.SetType(*i)
	}

	return grau
}

// AddType adds i to the "type" field.
func (grau *GameRoleAttributeUpdate) AddType(i int) *GameRoleAttributeUpdate {
	grau.mutation.AddType(i)
	return grau
}

// SetValue sets the "value" field.
func (grau *GameRoleAttributeUpdate) SetValue(s string) *GameRoleAttributeUpdate {
	grau.mutation.SetValue(s)
	return grau
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (grau *GameRoleAttributeUpdate) SetNillableValue(s *string) *GameRoleAttributeUpdate {
	if s != nil {
		grau.SetValue(*s)
	}

	return grau
}

// SetUpdateAt sets the "update_at" field.
func (grau *GameRoleAttributeUpdate) SetUpdateAt(t time.Time) *GameRoleAttributeUpdate {
	grau.mutation.SetUpdateAt(t)
	return grau
}

// Mutation returns the GameRoleAttributeMutation object of the builder.
func (grau *GameRoleAttributeUpdate) Mutation() *GameRoleAttributeMutation {
	return grau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (grau *GameRoleAttributeUpdate) Save(ctx context.Context) (int, error) {
	grau.defaults()
	return withHooks[int, GameRoleAttributeMutation](ctx, grau.sqlSave, grau.mutation, grau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (grau *GameRoleAttributeUpdate) SaveX(ctx context.Context) int {
	affected, err := grau.Save(ctx)
	if err != nil {
		panic(err)
	}

	return affected
}

// Exec executes the query.
func (grau *GameRoleAttributeUpdate) Exec(ctx context.Context) error {
	_, err := grau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grau *GameRoleAttributeUpdate) ExecX(ctx context.Context) {
	if err := grau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (grau *GameRoleAttributeUpdate) defaults() {
	if _, ok := grau.mutation.UpdateAt(); !ok {
		v := gameroleattribute.UpdateDefaultUpdateAt()
		grau.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (grau *GameRoleAttributeUpdate) check() error {
	if v, ok := grau.mutation.AccountID(); ok {
		if err := gameroleattribute.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`entity: validator failed for field "GameRoleAttribute.account_id": %w`, err)}
		}
	}

	if v, ok := grau.mutation.RoleID(); ok {
		if err := gameroleattribute.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`entity: validator failed for field "GameRoleAttribute.role_id": %w`, err)}
		}
	}

	if v, ok := grau.mutation.Name(); ok {
		if err := gameroleattribute.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entity: validator failed for field "GameRoleAttribute.name": %w`, err)}
		}
	}

	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (grau *GameRoleAttributeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GameRoleAttributeUpdate {
	grau.modifiers = append(grau.modifiers, modifiers...)
	return grau
}

func (grau *GameRoleAttributeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := grau.check(); err != nil {
		return n, err
	}

	_spec := sqlgraph.NewUpdateSpec(gameroleattribute.Table, gameroleattribute.Columns, sqlgraph.NewFieldSpec(gameroleattribute.FieldID, field.TypeInt64))
	if ps := grau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}

	if value, ok := grau.mutation.UserID(); ok {
		_spec.SetField(gameroleattribute.FieldUserID, field.TypeInt64, value)
	}

	if value, ok := grau.mutation.AddedUserID(); ok {
		_spec.AddField(gameroleattribute.FieldUserID, field.TypeInt64, value)
	}

	if value, ok := grau.mutation.AccountID(); ok {
		_spec.SetField(gameroleattribute.FieldAccountID, field.TypeString, value)
	}

	if value, ok := grau.mutation.RoleID(); ok {
		_spec.SetField(gameroleattribute.FieldRoleID, field.TypeString, value)
	}

	if value, ok := grau.mutation.Name(); ok {
		_spec.SetField(gameroleattribute.FieldName, field.TypeString, value)
	}

	if value, ok := grau.mutation.GetType(); ok {
		_spec.SetField(gameroleattribute.FieldType, field.TypeInt, value)
	}

	if value, ok := grau.mutation.AddedType(); ok {
		_spec.AddField(gameroleattribute.FieldType, field.TypeInt, value)
	}

	if value, ok := grau.mutation.Value(); ok {
		_spec.SetField(gameroleattribute.FieldValue, field.TypeString, value)
	}

	if value, ok := grau.mutation.UpdateAt(); ok {
		_spec.SetField(gameroleattribute.FieldUpdateAt, field.TypeTime, value)
	}

	_spec.AddModifiers(grau.modifiers...)

	if n, err = sqlgraph.UpdateNodes(ctx, grau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameroleattribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}

		return 0, err
	}

	grau.mutation.done = true

	return n, nil
}

// GameRoleAttributeUpdateOne is the builder for updating a single GameRoleAttribute entity.
type GameRoleAttributeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GameRoleAttributeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUserID sets the "user_id" field.
func (grauo *GameRoleAttributeUpdateOne) SetUserID(i int64) *GameRoleAttributeUpdateOne {
	grauo.mutation.ResetUserID()
	grauo.mutation.SetUserID(i)

	return grauo
}

// AddUserID adds i to the "user_id" field.
func (grauo *GameRoleAttributeUpdateOne) AddUserID(i int64) *GameRoleAttributeUpdateOne {
	grauo.mutation.AddUserID(i)
	return grauo
}

// SetAccountID sets the "account_id" field.
func (grauo *GameRoleAttributeUpdateOne) SetAccountID(s string) *GameRoleAttributeUpdateOne {
	grauo.mutation.SetAccountID(s)
	return grauo
}

// SetRoleID sets the "role_id" field.
func (grauo *GameRoleAttributeUpdateOne) SetRoleID(s string) *GameRoleAttributeUpdateOne {
	grauo.mutation.SetRoleID(s)
	return grauo
}

// SetName sets the "name" field.
func (grauo *GameRoleAttributeUpdateOne) SetName(s string) *GameRoleAttributeUpdateOne {
	grauo.mutation.SetName(s)
	return grauo
}

// SetType sets the "type" field.
func (grauo *GameRoleAttributeUpdateOne) SetType(i int) *GameRoleAttributeUpdateOne {
	grauo.mutation.ResetType()
	grauo.mutation.SetType(i)

	return grauo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (grauo *GameRoleAttributeUpdateOne) SetNillableType(i *int) *GameRoleAttributeUpdateOne {
	if i != nil {
		grauo.SetType(*i)
	}

	return grauo
}

// AddType adds i to the "type" field.
func (grauo *GameRoleAttributeUpdateOne) AddType(i int) *GameRoleAttributeUpdateOne {
	grauo.mutation.AddType(i)
	return grauo
}

// SetValue sets the "value" field.
func (grauo *GameRoleAttributeUpdateOne) SetValue(s string) *GameRoleAttributeUpdateOne {
	grauo.mutation.SetValue(s)
	return grauo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (grauo *GameRoleAttributeUpdateOne) SetNillableValue(s *string) *GameRoleAttributeUpdateOne {
	if s != nil {
		grauo.SetValue(*s)
	}

	return grauo
}

// SetUpdateAt sets the "update_at" field.
func (grauo *GameRoleAttributeUpdateOne) SetUpdateAt(t time.Time) *GameRoleAttributeUpdateOne {
	grauo.mutation.SetUpdateAt(t)
	return grauo
}

// Mutation returns the GameRoleAttributeMutation object of the builder.
func (grauo *GameRoleAttributeUpdateOne) Mutation() *GameRoleAttributeMutation {
	return grauo.mutation
}

// Where appends a list predicates to the GameRoleAttributeUpdate builder.
func (grauo *GameRoleAttributeUpdateOne) Where(ps ...predicate.GameRoleAttribute) *GameRoleAttributeUpdateOne {
	grauo.mutation.Where(ps...)
	return grauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (grauo *GameRoleAttributeUpdateOne) Select(field string, fields ...string) *GameRoleAttributeUpdateOne {
	grauo.fields = append([]string{field}, fields...)
	return grauo
}

// Save executes the query and returns the updated GameRoleAttribute entity.
func (grauo *GameRoleAttributeUpdateOne) Save(ctx context.Context) (*GameRoleAttribute, error) {
	grauo.defaults()
	return withHooks[*GameRoleAttribute, GameRoleAttributeMutation](ctx, grauo.sqlSave, grauo.mutation, grauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (grauo *GameRoleAttributeUpdateOne) SaveX(ctx context.Context) *GameRoleAttribute {
	node, err := grauo.Save(ctx)
	if err != nil {
		panic(err)
	}

	return node
}

// Exec executes the query on the entity.
func (grauo *GameRoleAttributeUpdateOne) Exec(ctx context.Context) error {
	_, err := grauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grauo *GameRoleAttributeUpdateOne) ExecX(ctx context.Context) {
	if err := grauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (grauo *GameRoleAttributeUpdateOne) defaults() {
	if _, ok := grauo.mutation.UpdateAt(); !ok {
		v := gameroleattribute.UpdateDefaultUpdateAt()
		grauo.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (grauo *GameRoleAttributeUpdateOne) check() error {
	if v, ok := grauo.mutation.AccountID(); ok {
		if err := gameroleattribute.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`entity: validator failed for field "GameRoleAttribute.account_id": %w`, err)}
		}
	}

	if v, ok := grauo.mutation.RoleID(); ok {
		if err := gameroleattribute.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`entity: validator failed for field "GameRoleAttribute.role_id": %w`, err)}
		}
	}

	if v, ok := grauo.mutation.Name(); ok {
		if err := gameroleattribute.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entity: validator failed for field "GameRoleAttribute.name": %w`, err)}
		}
	}

	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (grauo *GameRoleAttributeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GameRoleAttributeUpdateOne {
	grauo.modifiers = append(grauo.modifiers, modifiers...)
	return grauo
}

func (grauo *GameRoleAttributeUpdateOne) sqlSave(ctx context.Context) (_node *GameRoleAttribute, err error) {
	if err := grauo.check(); err != nil {
		return _node, err
	}

	_spec := sqlgraph.NewUpdateSpec(gameroleattribute.Table, gameroleattribute.Columns, sqlgraph.NewFieldSpec(gameroleattribute.FieldID, field.TypeInt64))

	id, ok := grauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entity: missing "GameRoleAttribute.id" for update`)}
	}

	_spec.Node.ID.Value = id
	if fields := grauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gameroleattribute.FieldID)

		for _, f := range fields {
			if !gameroleattribute.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entity: invalid field %q for query", f)}
			}

			if f != gameroleattribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}

	if ps := grauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}

	if value, ok := grauo.mutation.UserID(); ok {
		_spec.SetField(gameroleattribute.FieldUserID, field.TypeInt64, value)
	}

	if value, ok := grauo.mutation.AddedUserID(); ok {
		_spec.AddField(gameroleattribute.FieldUserID, field.TypeInt64, value)
	}

	if value, ok := grauo.mutation.AccountID(); ok {
		_spec.SetField(gameroleattribute.FieldAccountID, field.TypeString, value)
	}

	if value, ok := grauo.mutation.RoleID(); ok {
		_spec.SetField(gameroleattribute.FieldRoleID, field.TypeString, value)
	}

	if value, ok := grauo.mutation.Name(); ok {
		_spec.SetField(gameroleattribute.FieldName, field.TypeString, value)
	}

	if value, ok := grauo.mutation.GetType(); ok {
		_spec.SetField(gameroleattribute.FieldType, field.TypeInt, value)
	}

	if value, ok := grauo.mutation.AddedType(); ok {
		_spec.AddField(gameroleattribute.FieldType, field.TypeInt, value)
	}

	if value, ok := grauo.mutation.Value(); ok {
		_spec.SetField(gameroleattribute.FieldValue, field.TypeString, value)
	}

	if value, ok := grauo.mutation.UpdateAt(); ok {
		_spec.SetField(gameroleattribute.FieldUpdateAt, field.TypeTime, value)
	}

	_spec.AddModifiers(grauo.modifiers...)
	_node = &GameRoleAttribute{config: grauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues

	if err = sqlgraph.UpdateNode(ctx, grauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameroleattribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}

		return nil, err
	}

	grauo.mutation.done = true

	return _node, nil
}
