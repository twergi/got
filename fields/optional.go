package fields

// Optional - field that is optional. Might represent nullable fields,
// possible actions of specific objects attributes and etc.
type Optional[T any] struct {
	Value T `json:"value,omitempty"`
	// OK - is Value ok in any meaning
	OK bool `json:"is_valid,omitempty"`
}

// Ptr returns pointer to the value
func (o Optional[T]) Ptr() *T {
	if !o.OK {
		return nil
	}

	return &o.Value
}

// FromPtr returns Optional from pointer
func FromPtr[T any](v *T) Optional[T] {
	if v == nil {
		return Optional[T]{}
	}

	return New(*v)
}

// New returns Optional that is OK
func New[T any](v T) Optional[T] {
	return Optional[T]{
		Value: v,
		OK:    true,
	}
}

// FromMap returns Optional that is OK if map m has key k
func FromMap[T any, K comparable](m map[K]T, k K) Optional[T] {
	v, ok := m[k]
	if !ok {
		return Optional[T]{}
	}

	return New(v)
}

// FromOK ...
func FromOK[T any](v T, ok bool) Optional[T] {
	if !ok {
		return Optional[T]{}
	}
	return New(v)
}
