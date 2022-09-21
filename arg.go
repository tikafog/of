package of

import (
	"reflect"
)

type arg[T any] struct {
	t reflect.Type
	v T
}

//func (a Arg[T]) AssignAble(t reflect.Type) bool {
//	return a.t.AssignableTo(t)
//}

func (a arg[T]) Type() reflect.Type {
	return a.t
}

func (a arg[T]) Value() any {
	return a.v
}

type Arg interface {
	Type() reflect.Type
	Value() any
}

func NewArg[T any](v T) Arg {
	return &arg[T]{
		v: v,
		t: reflect.TypeOf(v),
	}
}

func ParseArgs[T any](args ...Arg) (T, bool) {
	for i := range args {

		if args[i].Type().AssignableTo(reflect.TypeOf(*new(T))) {
			return args[i].Value().(T), true
		}
	}
	return *new(T), false
}
