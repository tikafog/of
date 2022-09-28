package content

import (
	"encoding/json"

	"github.com/tikafog/of/buffers/content"
)

// ContentV3 ContentV3
type ContentV3 struct {
	meta    *metaContent
	From    string
	Message json.RawMessage
	Exts    []Ext
	Type    content.Type
}
