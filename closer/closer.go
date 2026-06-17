package closer

import (
	"slices"
	"sync"
)

type CloseFunc func()

var (
	mu      = &sync.Mutex{}
	closers []CloseFunc
)

func Bind(fn CloseFunc) {
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
