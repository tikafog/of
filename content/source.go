package content

import (
	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/feature/source"
)

func SourceType(p Type) source.Type {
	switch p {
	case content.TypeUpdate:
		return source.TypeUpgrade
	case content.TypeBootstrap:
		return source.TypeBootnode
	}
	return source.TypeMedia
}
