package of

import (
	"reflect"
)

type Arg interface {
	Type() string
	Value() interface{}
}

func ParseArgs[T any](args ...Arg) (T, bool) {
	for i := range args {
		if args[i].Type() == reflect.TypeOf(*new(T)).Name() {
			return args[i].Value().(T), true
		}
	}
	return *new(T), false
}
