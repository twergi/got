package lazy

import "context"

type LazyErr[T any] struct {
	lazy *LazyErrCtx[T]
}

func NewErr[T any](init func() (T, error)) *LazyErr[T] {
	return &LazyErr[T]{
		lazy: NewErrCtx(func(_ context.Context) (T, error) {
			return init()
		}),
	}
}

func (l *LazyErr[T]) Value() (T, error) {
	return l.lazy.Value(context.Background())
}
