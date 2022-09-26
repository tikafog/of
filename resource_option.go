package of

type ResourceOption func(*ResourceSetting)

type ResourceSetting struct {
	Type ResourceType
	//Path string
}

func SetResourceType(ResourceType ResourceType) ResourceOption {
	return func(option *ResourceSetting) {
		option.Type = ResourceType
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
