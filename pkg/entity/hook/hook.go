// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/lostsnow/keqing/pkg/entity"
)

// The ChatFunc type is an adapter to allow the use of ordinary
// function as Chat mutator.
type ChatFunc func(context.Context, *entity.ChatMutation) (entity.Value, error)

// Mutate calls f(ctx, m).
func (f ChatFunc) Mutate(ctx context.Context, m entity.Mutation) (entity.Value, error) {
	if mv, ok := m.(*entity.ChatMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entity.ChatMutation", m)
}

// The ChatOptionFunc type is an adapter to allow the use of ordinary
// function as ChatOption mutator.
type ChatOptionFunc func(context.Context, *entity.ChatOptionMutation) (entity.Value, error)

// Mutate calls f(ctx, m).
func (f ChatOptionFunc) Mutate(ctx context.Context, m entity.Mutation) (entity.Value, error) {
	if mv, ok := m.(*entity.ChatOptionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entity.ChatOptionMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *entity.UserMutation) (entity.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m entity.Mutation) (entity.Value, error) {
	if mv, ok := m.(*entity.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entity.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, entity.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entity.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entity.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m entity.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op entity.Op) Condition {
	return func(_ context.Context, m entity.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entity.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entity.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entity.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk entity.Hook, cond Condition) entity.Hook {
	return func(next entity.Mutator) entity.Mutator {
		return entity.MutateFunc(func(ctx context.Context, m entity.Mutation) (entity.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, entity.Delete|entity.Create)
func On(hk entity.Hook, op entity.Op) entity.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, entity.Update|entity.UpdateOne)
func Unless(hk entity.Hook, op entity.Op) entity.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) entity.Hook {
	return func(entity.Mutator) entity.Mutator {
		return entity.MutateFunc(func(context.Context, entity.Mutation) (entity.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []entity.Hook {
//		return []entity.Hook{
//			Reject(entity.Delete|entity.Update),
//		}
//	}
func Reject(op entity.Op) entity.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []entity.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...entity.Hook) Chain {
	return Chain{append([]entity.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() entity.Hook {
	return func(mutator entity.Mutator) entity.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...entity.Hook) Chain {
	newHooks := make([]entity.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
