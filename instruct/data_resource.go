package instruct

const CurrentResourceVersion = 1

type ResourceData struct {
	Version int      `json:"version,omitempty"`
	List    []string `json:"list,omitempty"`
	Action  int      `json:"action,omitempty"`
	Last    int64    `json:"last,omitempty"`
}

func NewResourceData() *ResourceData {
	return &ResourceData{
		Version: CurrentResourceVersion,
	}
}
