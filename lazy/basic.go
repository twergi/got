package lazy

import (
	"context"
	"sync"
)

type Lazy[T any] struct {
	val  T
	init func(context.Context) T
	once sync.Once
}

func NewCtx[T any](init func(context.Context) T) *Lazy[T] {
	return &Lazy[T]{
		init: init,
	}
}

func New[T any](init func() T) *Lazy[T] {
	return NewCtx(func(ctx context.Context) T { return init() })
}

func (l *Lazy[T]) Value(ctx context.Context) T {
	l.once.Do(func() { l.val = l.init(ctx) })
	return l.val
}
