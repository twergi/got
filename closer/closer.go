package closer

import (
	"slices"
	"sync"
)

var (
	mu      = &sync.Mutex{}
	closers []func()
)

func Bind(fn func()) {
	mu.Lock()
	defer mu.Unlock()
	closers = append(closers, fn)
}

func Close() {
	mu.Lock()
	defer mu.Unlock()
	for _, fn := range slices.Backward(closers) {
		fn()
	}
	closers = nil
}
