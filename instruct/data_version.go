package instruct

var CurrentDataVersion = DataVersion{Version: 1}

type DataVersion struct {
	Version int `json:"version,omitempty"`
}
