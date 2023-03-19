// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lostsnow/keqing/pkg/entity/gameaccount"
	"github.com/lostsnow/keqing/pkg/entity/predicate"
)

// GameAccountDelete is the builder for deleting a GameAccount entity.
type GameAccountDelete struct {
	config
	hooks    []Hook
	mutation *GameAccountMutation
}

// Where appends a list predicates to the GameAccountDelete builder.
func (gad *GameAccountDelete) Where(ps ...predicate.GameAccount) *GameAccountDelete {
	gad.mutation.Where(ps...)
	return gad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gad *GameAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GameAccountMutation](ctx, gad.sqlExec, gad.mutation, gad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gad *GameAccountDelete) ExecX(ctx context.Context) int {
	n, err := gad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gad *GameAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(gameaccount.Table, sqlgraph.NewFieldSpec(gameaccount.FieldID, field.TypeInt64))
	if ps := gad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gad.mutation.done = true
	return affected, err
}

// GameAccountDeleteOne is the builder for deleting a single GameAccount entity.
type GameAccountDeleteOne struct {
	gad *GameAccountDelete
}

// Where appends a list predicates to the GameAccountDelete builder.
func (gado *GameAccountDeleteOne) Where(ps ...predicate.GameAccount) *GameAccountDeleteOne {
	gado.gad.mutation.Where(ps...)
	return gado
}

// Exec executes the deletion query.
func (gado *GameAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := gado.gad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gameaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gado *GameAccountDeleteOne) ExecX(ctx context.Context) {
	if err := gado.Exec(ctx); err != nil {
		panic(err)
	}
}
