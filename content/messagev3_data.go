package content

//MessageV3Data ...
//@Description:
type MessageV3Data[T any] struct {
	Last    int64 `json:"last,omitempty"`
	Index   int64 `json:"index,omitempty"`
	Version int   `json:"version,omitempty"` // current info version
	Length  int   `json:"length,omitempty"`
	Data    []*T  `json:"data,omitempty"`
}
