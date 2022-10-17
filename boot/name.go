package boot

import (
	"github.com/tikafog/of"
)

func LoadName(id uint64) of.Name {
	if v, ok := names.Get(id); ok {
		return v
	}
	return of.NameNotSet
}
