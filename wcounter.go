package wcounter

import (
	"sync"
	"time"
)

type WindowCounter[T comparable] struct {
	items  map[T]int
	window time.Duration
	mutex  sync.RWMutex
}

func New[T comparable](window time.Duration) *WindowCounter[T] {
	return &WindowCounter[T]{
		items:  make(map[T]int),
		window: window,
	}
}

func (wc *WindowCounter[T]) Add(item T) {
	wc.mutex.Lock()
	_, found := wc.items[item]
	if !found {
		wc.items[item] = 1
		go wc.clear(item)
	} else {
		wc.items[item]++
	}
	wc.mutex.Unlock()
}

func (wc *WindowCounter[T]) Get(item T) int {
	wc.mutex.RLock()
	defer wc.mutex.RUnlock()
	return wc.items[item]
}

func (wc *WindowCounter[T]) clear(item T) {
	time.Sleep(wc.window)
	wc.mutex.Lock()
	delete(wc.items, item)
	wc.mutex.Unlock()
}
