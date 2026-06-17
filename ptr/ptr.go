package ptr

func New[T any](v T) *T {
	return new(v)
}

func Deref[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}

	return *v
}
