package wcounter

import (
	"testing"
	"time"
)

func TestFlow(t *testing.T) {
	wc := New[string](1 * time.Second)
	item := "test"

	counter := wc.Get(item)
	if counter != 0 {
		t.Fatalf("Didn't return 0: counter=%d", counter)
	}

	wc.Add(item)

	counter = wc.Get(item)
	if counter != 1 {
		t.Fatalf("Didn't return 1: counter=%d", counter)
	}

	wc.Add(item)

	counter = wc.Get(item)
	if counter != 2 {
		t.Fatalf("Didn't return 2: counter=%d", counter)
	}

	go wc.Add(item)
	go wc.Add(item)
	go wc.Add(item)

	time.Sleep(10 * time.Millisecond)

	counter = wc.Get(item)
	if counter != 5 {
		t.Fatalf("Didn't return 5: counter=%d", counter)
	}

	time.Sleep(1500 * time.Millisecond)

	counter = wc.Get(item)
	if counter != 0 {
		t.Fatalf("Didn't return 0: counter=%d", counter)
	}
}

func TestFlowWithStruct(t *testing.T) {
	type TestData struct {
		Text   string
		Number int
	}

	wc := New[TestData](1 * time.Second)
	item := TestData{
		Text:   t.Name(),
		Number: 42,
	}

	counter := wc.Get(item)
	if counter != 0 {
		t.Fatalf("Didn't return 0: counter=%d", counter)
	}

	wc.Add(item)

	counter = wc.Get(item)
	if counter != 1 {
		t.Fatalf("Didn't return 1: counter=%d", counter)
	}

	wc.Add(item)

	counter = wc.Get(item)
	if counter != 2 {
		t.Fatalf("Didn't return 2: counter=%d", counter)
	}

	go wc.Add(item)
	go wc.Add(item)
	go wc.Add(item)

	time.Sleep(10 * time.Millisecond)

	counter = wc.Get(item)
	if counter != 5 {
		t.Fatalf("Didn't return 5: counter=%d", counter)
	}

	time.Sleep(1500 * time.Millisecond)

	counter = wc.Get(item)
	if counter != 0 {
		t.Fatalf("Didn't return 0: counter=%d", counter)
	}
}
