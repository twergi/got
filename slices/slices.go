package slices

func Map[T, K any](vs []T, f func(T) K) []K {
	ret := make([]K, 0, len(vs))
	for _, v := range vs {
		ret = append(ret, f(v))
	}
	return ret
}

func MapToMap[T any, K comparable](vs []T, getKey func(T) K) map[K]T {
	m := make(map[K]T, len(vs))

	for _, v := range vs {
		m[getKey(v)] = v
	}

	return m
}

func Group[T any, K comparable](vs []T, getKey func(T) K) map[K][]T {
	m := make(map[K][]T)

	for _, v := range vs {
		k := getKey(v)
		m[k] = append(m[k], v)
	}

	return m
}

func Filter[T any](vs []T, filterFunc func(T) bool) []T {
	r := make([]T, 0)
	for _, v := range vs {
		if !filterFunc(v) {
			continue
		}

		r = append(r, v)
	}

	return r
}

func Unique[T comparable](vs []T) []T {
	s := make(map[T]struct{}, len(vs))
	r := make([]T, 0)

	for _, v := range vs {
		if _, dup := s[v]; dup {
			continue
		}
		s[v] = struct{}{}

		r = append(r, v)
	}

	return r
}

func Flatten[T any](vs [][]T) []T {
	size := 0
	for _, v := range vs {
		size += len(v)
	}

	r := make([]T, 0, size)
	for _, v := range vs {
		r = append(r, v...)
	}

	return r
}

func Zip[T any](v1, v2 []T) []T {
	minItems := min(len(v1), len(v2))
	ret := make([]T, 0, 2*minItems)
	for i := range minItems {
		ret = append(ret, v1[i], v2[i])
	}

	return ret
}

func Slice[T any](v []T, n int) [][]T {
	if len(v) == 0 {
		return nil
	}

	if n < 1 {
		return nil
	}

	size := len(v) / n

	// 4 / 3 = 1
	if size*n != len(v) {
		size++
	}

	res := make([][]T, 0, size)
	for i := range size {
		res = append(res, make([]T, 0, n))
		k := i * n
		for j := range n {
			if k+j == len(v) {
				break
			}

			res[i] = append(res[i], v[k+j])
		}
	}
	return res
}
