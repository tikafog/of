package resource

type AddOption func(*AddSetting)

type AddSetting struct {
	Type         Type
	ShowProgress bool
	IgnoreHidden bool
}

// TypeOption ...
// @param ResourceType
// @return ResourceOption
func TypeOption(ResourceType Type) AddOption {
	return func(option *AddSetting) {
		option.Type = ResourceType
	}
}

// ShowProgressOption ...
// @param bool
// @return ResourceOption
func ShowProgressOption() AddOption {
	return func(option *AddSetting) {
		option.ShowProgress = true
	}
}

// IgnoreHiddenOption ...
// @param bool
// @return ResourceOption
func IgnoreHiddenOption() AddOption {
	return func(option *AddSetting) {
		option.IgnoreHidden = true
	}
}
