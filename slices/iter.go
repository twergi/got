package slices

import (
	"iter"
)

func MapSeq[T, K any](iter iter.Seq[T], mapFunc func(T) K) iter.Seq[K] {
	return func(yield func(K) bool) {
		for v := range iter {
			if !yield(mapFunc(v)) {
				return
			}
		}
	}
}

func UniqueSeq[T comparable](iter iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		s := make(map[T]struct{})
		for v := range iter {
			if _, dup := s[v]; dup {
				continue
			}
			s[v] = struct{}{}

			if !yield(v) {
				return
			}
		}
	}
}

func FilterSeq[T any](iter iter.Seq[T], filterFunc func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range iter {
			if !filterFunc(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

func FlattenSeq[T any](iter iter.Seq[[]T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range iter {
			for _, vv := range v {
				if !yield(vv) {
					return
				}
			}
		}
	}
}

func ZipSeq[T any](v1, v2 []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		minItems := min(len(v1), len(v2))
		for i := range minItems {
			if !yield(v1[i]) || !yield(v2[i]) {
				return
			}
		}
	}
}

func SliceSeq[T any](iter iter.Seq[T], n int) iter.Seq[[]T] {
	if n < 1 {
		return nil
	}

	return func(yield func([]T) bool) {
		batch := make([]T, 0, n)
		for v := range iter {
			batch = append(batch, v)
			if len(batch) != n {
				continue
			}
			if !yield(batch) {
				return
			}
			batch = make([]T, 0, n)
		}

		if len(batch) > 0 {
			yield(batch)
		}
	}
}
