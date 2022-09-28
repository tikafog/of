// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/tikafog/of/dbc/media"
)

// The AnnounceFunc type is an adapter to allow the use of ordinary
// function as Announce mutator.
type AnnounceFunc func(context.Context, *media.AnnounceMutation) (media.Value, error)

// Mutate calls f(ctx, m).
func (f AnnounceFunc) Mutate(ctx context.Context, m media.Mutation) (media.Value, error) {
	mv, ok := m.(*media.AnnounceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *media.AnnounceMutation", m)
	}
	return f(ctx, mv)
}

// The ChannelFunc type is an adapter to allow the use of ordinary
// function as Channel mutator.
type ChannelFunc func(context.Context, *media.ChannelMutation) (media.Value, error)

// Mutate calls f(ctx, m).
func (f ChannelFunc) Mutate(ctx context.Context, m media.Mutation) (media.Value, error) {
	mv, ok := m.(*media.ChannelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *media.ChannelMutation", m)
	}
	return f(ctx, mv)
}

// The DiscoveryFunc type is an adapter to allow the use of ordinary
// function as Discovery mutator.
type DiscoveryFunc func(context.Context, *media.DiscoveryMutation) (media.Value, error)

// Mutate calls f(ctx, m).
func (f DiscoveryFunc) Mutate(ctx context.Context, m media.Mutation) (media.Value, error) {
	mv, ok := m.(*media.DiscoveryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *media.DiscoveryMutation", m)
	}
	return f(ctx, mv)
}

// The InformationV1Func type is an adapter to allow the use of ordinary
// function as InformationV1 mutator.
type InformationV1Func func(context.Context, *media.InformationV1Mutation) (media.Value, error)

// Mutate calls f(ctx, m).
func (f InformationV1Func) Mutate(ctx context.Context, m media.Mutation) (media.Value, error) {
	mv, ok := m.(*media.InformationV1Mutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *media.InformationV1Mutation", m)
	}
	return f(ctx, mv)
}

// The PageFunc type is an adapter to allow the use of ordinary
// function as Page mutator.
type PageFunc func(context.Context, *media.PageMutation) (media.Value, error)

// Mutate calls f(ctx, m).
func (f PageFunc) Mutate(ctx context.Context, m media.Mutation) (media.Value, error) {
	mv, ok := m.(*media.PageMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *media.PageMutation", m)
	}
	return f(ctx, mv)
}

// The TopListFunc type is an adapter to allow the use of ordinary
// function as TopList mutator.
type TopListFunc func(context.Context, *media.TopListMutation) (media.Value, error)

// Mutate calls f(ctx, m).
func (f TopListFunc) Mutate(ctx context.Context, m media.Mutation) (media.Value, error) {
	mv, ok := m.(*media.TopListMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *media.TopListMutation", m)
	}
	return f(ctx, mv)
}

// The VersionFunc type is an adapter to allow the use of ordinary
// function as Version mutator.
type VersionFunc func(context.Context, *media.VersionMutation) (media.Value, error)

// Mutate calls f(ctx, m).
func (f VersionFunc) Mutate(ctx context.Context, m media.Mutation) (media.Value, error) {
	mv, ok := m.(*media.VersionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *media.VersionMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, media.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m media.Mutation) bool {
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
	return func(ctx context.Context, m media.Mutation) bool {
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
	return func(ctx context.Context, m media.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op media.Op) Condition {
	return func(_ context.Context, m media.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m media.Mutation) bool {
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
	return func(_ context.Context, m media.Mutation) bool {
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
	return func(_ context.Context, m media.Mutation) bool {
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
func If(hk media.Hook, cond Condition) media.Hook {
	return func(next media.Mutator) media.Mutator {
		return media.MutateFunc(func(ctx context.Context, m media.Mutation) (media.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, media.Delete|media.Create)
func On(hk media.Hook, op media.Op) media.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, media.Update|media.UpdateOne)
func Unless(hk media.Hook, op media.Op) media.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) media.Hook {
	return func(media.Mutator) media.Mutator {
		return media.MutateFunc(func(context.Context, media.Mutation) (media.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []media.Hook {
//		return []media.Hook{
//			Reject(media.Delete|media.Update),
//		}
//	}
func Reject(op media.Op) media.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []media.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...media.Hook) Chain {
	return Chain{append([]media.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() media.Hook {
	return func(mutator media.Mutator) media.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...media.Hook) Chain {
	newHooks := make([]media.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
