// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/tikafog/of/dbc/bootnode"
)

// The BootstrapFunc type is an adapter to allow the use of ordinary
// function as Bootstrap mutator.
type BootstrapFunc func(context.Context, *bootnode.BootstrapMutation) (bootnode.Value, error)

// Mutate calls f(ctx, m).
func (f BootstrapFunc) Mutate(ctx context.Context, m bootnode.Mutation) (bootnode.Value, error) {
	mv, ok := m.(*bootnode.BootstrapMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *bootnode.BootstrapMutation", m)
	}
	return f(ctx, mv)
}

// The VersionFunc type is an adapter to allow the use of ordinary
// function as Version mutator.
type VersionFunc func(context.Context, *bootnode.VersionMutation) (bootnode.Value, error)

// Mutate calls f(ctx, m).
func (f VersionFunc) Mutate(ctx context.Context, m bootnode.Mutation) (bootnode.Value, error) {
	mv, ok := m.(*bootnode.VersionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *bootnode.VersionMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, bootnode.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m bootnode.Mutation) bool {
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
	return func(ctx context.Context, m bootnode.Mutation) bool {
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
	return func(ctx context.Context, m bootnode.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op bootnode.Op) Condition {
	return func(_ context.Context, m bootnode.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m bootnode.Mutation) bool {
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
	return func(_ context.Context, m bootnode.Mutation) bool {
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
	return func(_ context.Context, m bootnode.Mutation) bool {
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
//
func If(hk bootnode.Hook, cond Condition) bootnode.Hook {
	return func(next bootnode.Mutator) bootnode.Mutator {
		return bootnode.MutateFunc(func(ctx context.Context, m bootnode.Mutation) (bootnode.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, bootnode.Delete|bootnode.Create)
//
func On(hk bootnode.Hook, op bootnode.Op) bootnode.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, bootnode.Update|bootnode.UpdateOne)
//
func Unless(hk bootnode.Hook, op bootnode.Op) bootnode.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) bootnode.Hook {
	return func(bootnode.Mutator) bootnode.Mutator {
		return bootnode.MutateFunc(func(context.Context, bootnode.Mutation) (bootnode.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []bootnode.Hook {
//		return []bootnode.Hook{
//			Reject(bootnode.Delete|bootnode.Update),
//		}
//	}
//
func Reject(op bootnode.Op) bootnode.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []bootnode.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...bootnode.Hook) Chain {
	return Chain{append([]bootnode.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() bootnode.Hook {
	return func(mutator bootnode.Mutator) bootnode.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...bootnode.Hook) Chain {
	newHooks := make([]bootnode.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}