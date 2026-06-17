package set

import (
	"iter"
	"maps"
)

type Set[T comparable] struct {
	v map[T]struct{}
}

func New[T comparable](size ...int) Set[T] {
	if len(size) == 0 {
		return Set[T]{
			v: make(map[T]struct{}),
		}
	}

	if len(size) > 1 {
		panic("expected 1 value in size")
	}

	return Set[T]{
		v: make(map[T]struct{}, size[0]),
	}
}

func Of[T comparable](vs ...T) Set[T] {
	s := New[T](len(vs))
	for _, v := range vs {
		s.Put(v)
	}
	return s
}

func Collect[T comparable](vs iter.Seq[T]) Set[T] {
	s := New[T]()
	for v := range vs {
		s.Put(v)
	}
	return s
}

func (s Set[T]) Put(v T) {
	s.v[v] = struct{}{}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s.v[v]
	return ok
}

func (s Set[T]) Remove(v T) {
	delete(s.v, v)
}

func (s Set[T]) Size() int {
	return len(s.v)
}

func (s Set[T]) Each(f func(k T) bool) {
	for k := range s.v {
		if !f(k) {
			break
		}
	}
}

func (s Set[T]) Collect(vs iter.Seq[T]) {
	for v := range vs {
		s.Put(v)
	}
}

func (s Set[T]) GetElements() map[T]struct{} {
	return s.v
}

func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s.v)
}

func (s Set[T]) Slice() []T {
	slice := make([]T, 0, len(s.v))
	for k := range s.v {
		slice = append(slice, k)
	}
	return slice
}

func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	intersection := New[T]()
	for k := range s.v {
		if s2.Has(k) {
			intersection.Put(k)
		}
	}
	return intersection
}

func (s Set[T]) Unite(s2 Set[T]) Set[T] {
	union := New[T]()
	for k := range s.v {
		union.Put(k)
	}
	for k := range s2.v {
		union.Put(k)
	}
	return union
}

func (s Set[T]) SymmetricalDifference(s2 Set[T]) Set[T] {
	diff := Of(s.Slice()...)
	for k := range s2.v {
		if diff.Has(k) {
			diff.Remove(k)
		} else {
			diff.Put(k)
		}
	}
	return diff
}
