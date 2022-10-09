package of

type ResourceOption func(*ResourceSetting)

type ResourceSetting struct {
	Type         ResourceType
	ShowProgress bool
	IgnoreHidden bool
}

// SetResourceType ...
// @param ResourceType
// @return ResourceOption
func SetResourceType(ResourceType ResourceType) ResourceOption {
	return func(option *ResourceSetting) {
		option.Type = ResourceType
	}
}

// SetResourceShowProgress ...
// @param bool
// @return ResourceOption
func SetResourceShowProgress(ShowProgress bool) ResourceOption {
	return func(option *ResourceSetting) {
		option.ShowProgress = ShowProgress
	}
}

//func SetResourcePath(path string) ResourceOption {
//	return func(option *ResourceSetting) {
//		option.Path = path
//	}
//}
//
//func SetResource(path string, ResourceType ResourceType) ResourceOption {
//	return func(option *ResourceSetting) {
//		option.Type = ResourceType
//		option.Path = path
//	}
//}
