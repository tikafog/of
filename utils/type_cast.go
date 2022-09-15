package utils

func TypeCase[T any](v any) T {
	return v.(T)
}
