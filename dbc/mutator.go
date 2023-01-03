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
	if value == nil {
		return nil, merr.New("mutate: privacy(query) not supported update")
	}
	//v, ok := value.(sync.Locker)
	//if !ok {
	//	return nil, merr.New("mutate: privacy is not correctly implemented")
	//}
	//v.Lock()
	return m.parent.Mutate(ctx, mutation)
}
