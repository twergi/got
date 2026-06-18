package lazy

import "context"

type Lazy[T any] struct {
	lazy *LazyErrCtx[T]
}

func New[T any](init func() T) *Lazy[T] {
	return &Lazy[T]{
		lazy: NewErrCtx(func(ctx context.Context) (T, error) {
			return init(), nil
		}),
	}
}

func (l *Lazy[T]) Value() T {
	v, _ := l.lazy.Value(context.Background())
	return v
}
