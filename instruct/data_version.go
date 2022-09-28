package instruct

var CurrentDataVersion = VersionData{Version: 1}

type VersionData struct {
	Version int `json:"version,omitempty"`
}
