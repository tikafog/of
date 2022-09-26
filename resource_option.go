package of

type ResourceOption func(*resourceOption)

type resourceOption struct {
	Type ResourceType
}

func SetResourceType(ResourceType ResourceType) ResourceOption {
	return func(option *resourceOption) {
		option.Type = ResourceType
	}
}
