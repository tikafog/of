package resource

type AddSetting struct {
	Type         Type
	ShowProgress bool
	IgnoreHidden bool
}

type RemoveSetting struct {
	Type Type
}
