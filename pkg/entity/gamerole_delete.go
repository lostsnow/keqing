// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lostsnow/keqing/pkg/entity/gamerole"
	"github.com/lostsnow/keqing/pkg/entity/predicate"
)

// GameRoleDelete is the builder for deleting a GameRole entity.
type GameRoleDelete struct {
	config
	hooks    []Hook
	mutation *GameRoleMutation
}

// Where appends a list predicates to the GameRoleDelete builder.
func (grd *GameRoleDelete) Where(ps ...predicate.GameRole) *GameRoleDelete {
	grd.mutation.Where(ps...)
	return grd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (grd *GameRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GameRoleMutation](ctx, grd.sqlExec, grd.mutation, grd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (grd *GameRoleDelete) ExecX(ctx context.Context) int {
	n, err := grd.Exec(ctx)
	if err != nil {
		panic(err)
	}

	return n
}

func (grd *GameRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(gamerole.Table, sqlgraph.NewFieldSpec(gamerole.FieldID, field.TypeInt64))
	if ps := grd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}

	affected, err := sqlgraph.DeleteNodes(ctx, grd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}

	grd.mutation.done = true

	return affected, err
}

// GameRoleDeleteOne is the builder for deleting a single GameRole entity.
type GameRoleDeleteOne struct {
	grd *GameRoleDelete
}

// Where appends a list predicates to the GameRoleDelete builder.
func (grdo *GameRoleDeleteOne) Where(ps ...predicate.GameRole) *GameRoleDeleteOne {
	grdo.grd.mutation.Where(ps...)
	return grdo
}

// Exec executes the deletion query.
func (grdo *GameRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := grdo.grd.Exec(ctx)

	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gamerole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (grdo *GameRoleDeleteOne) ExecX(ctx context.Context) {
	if err := grdo.Exec(ctx); err != nil {
		panic(err)
	}
}
