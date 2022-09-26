package of

import (
	"context"
)

//ResourceType ...
//ENUM(local,remote,hash,remote hash,max)
type ResourceType uint32

type Resource interface {
	Add(ctx context.Context, path string, ro ...ResourceOption) error
	Ls(ctx context.Context, path string, ro ...ResourceOption) error
	Remove(ctx context.Context, path string, ro ...ResourceOption) error
	Validate(ctx context.Context, path string, ro ...ResourceOption) error
}
