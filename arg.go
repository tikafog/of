package of

import (
	"reflect"
)

type Arg interface {
	Type() reflect.Type
	Value() any
}

func ParseArgs[T any](args ...Arg) (T, bool) {
	for i := range args {

		if args[i].Type().AssignableTo(reflect.TypeOf(*new(T))) {
			return args[i].Value().(T), true
		}
	}
	return *new(T), false
}
