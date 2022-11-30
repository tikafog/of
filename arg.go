package of

import (
	"reflect"
)

// KeyArg ...
type KeyArg struct {
	Key   any
	Value any
}

// Args ...
type Args map[any]any

// Get ...
// @receiver Args
// @param string
// @return any
// @return bool
func (a Args) Get(key any) (any, bool) {
	v, ok := a[key]
	return v, ok
}

// GetKeyArg ...
// @receiver Args
// @param string
// @return *KeyArg
// @return bool
func (a Args) GetKeyArg(key any) (*KeyArg, bool) {
	v, ok := a.Get(key)
	if !ok {
		return nil, ok
	}
	return &KeyArg{
		Key:   key,
		Value: v,
	}, true
}

// Add ...
// @receiver Args
// @param KeyArg
// @return any
// @return bool
func (a Args) Add(arg KeyArg) (any, bool) {
	if v, ok := a[arg.Key]; ok {
		return v, false
	}
	a[arg.Key] = arg.Value
	return arg.Value, true
}

// AddKeyArg ...
// @receiver Args
// @param string
// @param any
// @return any
// @return bool
func (a Args) AddKeyArg(key string, arg any) (any, bool) {
	if v, ok := a[key]; ok {
		return v, false
	}
	a[key] = arg
	return arg, true
}

// RangeArgs ...
// @receiver Args
// @param func(key any, arg any) bool
func (a Args) RangeArgs(fn func(key any, arg any) bool) {
	for k, v := range a {
		if !fn(k, v) {
			return
		}
	}
}

// NewArgs ...
// @param ...KeyArg
// @return Args
func NewArgs(args ...KeyArg) Args {
	v := make(map[any]any, len(args))
	for i := range args {
		v[args[i].Key] = args[i].Value
	}
	return v
}

func NewTypeArg(arg any) KeyArg {
	return KeyArg{
		Key:   reflect.TypeOf(arg),
		Value: arg,
	}
}

// NewKeyArg ...
// @param string
// @param any
// @return KeyArg
func NewKeyArg(key string, arg any) KeyArg {
	return KeyArg{
		Key:   key,
		Value: arg,
	}
}

// TypeGetArgs ...
// @param Args
// @return any
// @return bool
func TypeGetArgs[T any](args Args) (any, bool) {
	t := reflect.TypeOf(*new(T))
	return args.Get(t)
}
