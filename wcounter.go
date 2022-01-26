package wcounter

import (
	"sync"
	"time"
)

type WindowCounter struct {
	items  map[interface{}]int
	window time.Duration
	mutex  sync.RWMutex
}

func New(window time.Duration) *WindowCounter {
	return &WindowCounter{
		items:  make(map[interface{}]int),
		window: window,
	}
}

func (wc *WindowCounter) Add(item interface{}) {
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

func (wc *WindowCounter) Get(item interface{}) int {
	wc.mutex.RLock()
	defer wc.mutex.RUnlock()
	return wc.items[item]
}

func (wc *WindowCounter) clear(item interface{}) {
	time.Sleep(wc.window)
	wc.mutex.Lock()
	delete(wc.items, item)
	wc.mutex.Unlock()
}
