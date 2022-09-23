package instruct

type Resource struct {
	List   []string `json:"list,omitempty"`
	Action int      `json:"action,omitempty"`
	Last   int64    `json:"last,omitempty"`
}
