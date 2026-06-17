package lazy

import (
	"context"
	"sync"
)

type LazyErr[T any] struct {
	val  T
	init func(context.Context) (T, error)
	once sync.Once
	err  error
}

func NewErrCtx[T any](init func(context.Context) (T, error)) *LazyErr[T] {
	return &LazyErr[T]{
		init: init,
	}
}

func NewErr[T any](init func() (T, error)) *LazyErr[T] {
	return NewErrCtx(func(ctx context.Context) (T, error) {
		return init()
	})
}

func (l *LazyErr[T]) Value(ctx context.Context) (T, error) {
	l.once.Do(func() {
		l.val, l.err = l.init(ctx)
	})
	return l.val, l.err
}
