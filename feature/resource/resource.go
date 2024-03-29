package resource

import (
	"context"
)

// Type ...
// ENUM(local,remote,hash,remote_hash,max)
type Type uint32

type Resource interface {
	Add(ctx context.Context, path string, setting AddSetting) (string, error)
	//Ls(ctx context.Context, path string, ro ...ResourceOption) (string, error)
	Remove(ctx context.Context, path string, setting RemoveSetting) error
	//Validate(ctx context.Context, path string, ro ...ResourceOption) error
}
