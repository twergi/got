package lazy

import (
	"context"
)

type LazyCtx[T any] struct {
	lazy *LazyErrCtx[T]
}

func NewCtx[T any](init func(context.Context) T) *LazyCtx[T] {
	return &LazyCtx[T]{
		lazy: NewErrCtx(func(ctx context.Context) (T, error) {
			return init(ctx), nil
		}),
	}
}

func (l *LazyCtx[T]) Value(ctx context.Context) T {
	v, _ := l.lazy.Value(ctx)
	return v
}
