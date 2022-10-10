package dbc

import (
	"context"

	"entgo.io/ent"
)

type Mutator struct {
	parent ent.Mutator
}

var MutatorFunc = func(mutator ent.Mutator) ent.Mutator {
	return &Mutator{
		parent: mutator,
	}
}

func (m *Mutator) Mutate(ctx context.Context, mutation ent.Mutation) (ent.Value, error) {
	value := ctx.Value("privacy")
	if value == nil || value == "query" {
		return nil, merr.New("query not supported update")
	}
	return m.parent.Mutate(ctx, mutation)
}
