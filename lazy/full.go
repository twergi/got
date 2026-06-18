package lazy

import (
	"context"
	"sync"
)

type LazyErrCtx[T any] struct {
	val  T
	err  error
	init func(context.Context) (T, error)
	once sync.Once
}

func NewErrCtx[T any](init func(context.Context) (T, error)) *LazyErrCtx[T] {
	return &LazyErrCtx[T]{
		init: init,
	}
}

func (l *LazyErrCtx[T]) Value(ctx context.Context) (T, error) {
	l.once.Do(func() { l.val, l.err = l.init(ctx) })
	return l.val, l.err
}
