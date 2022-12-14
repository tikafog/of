package eload

import "github.com/tikafog/of"

type noSet struct{}

func (n noSet) Name() of.Name {
	return of.NameNotSet
}

func (n noSet) IsValid() bool {
	return false
}

var NoSet of.Module = &noSet{}
