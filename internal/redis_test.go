package redis

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestSetGet(t *testing.T) {
	store := NewStorage()

	keys := []string{"space", "int", "float", "name"}
	values := []string{"", "1", "1.0", "Anton"}

	for i := range keys {
		store.Set(keys[i], values[i])
	}

	for i := range keys {
		got, ok := store.Get(keys[i])
		if !ok {
			t.Fatalf("key %q not found", keys[i])
		}
		if got != values[i] {
			t.Errorf("key %q: got %q, want %q", keys[i], got, values[i])
		}
	}
}

func TestGetNotFound(t *testing.T) {
	store := NewStorage()

	fakeKey := "Aboba"
	keys := []string{"space", "int", "float", "name"}
	values := []string{"", "1", "1.0", "Anton"}

	for i := range keys {
		store.Set(keys[i], values[i])
	}

	_, ok := store.Get(fakeKey)
	if ok {
		t.Fatalf("key %q found", fakeKey)
	}
}

func TestDelete(t *testing.T) {
	store := NewStorage()

	keys := []string{"space", "int", "float", "name"}
	values := []string{"", "1", "1.0", "Anton"}

	for i := range keys {
		store.Set(keys[i], values[i])
	}

	_, exist := store.Get(keys[0])
	if !exist {
		t.Fatalf("key %q not set in init", keys[0])
	}

	store.Delete(keys[0])

	_, ok := store.Get(keys[0])
	if ok {
		t.Fatalf("key %q not delete", keys[0])
	}
}

func TestDeleteNotFound(t *testing.T) {
	store := NewStorage()

	fakeKey := "Aboba"
	keys := []string{"space", "int", "float", "name"}
	values := []string{"", "1", "1.0", "Anton"}

	for i := range keys {
		store.Set(keys[i], values[i])
	}

	_, ok := store.Get(fakeKey)
	if ok {
		t.Fatalf("fakekey %q found in storage.", fakeKey)
	}
	store.Delete(fakeKey)

	for i := range keys {
		got, ok := store.Get(keys[i])
		if !ok {
			t.Fatalf("key %q not found", keys[i])
		}
		if got != values[i] {
			t.Errorf("key %q: got %q, want %q", keys[i], got, values[i])
		}
	}
}

func TestParallel(t *testing.T) {
	store := NewStorage()
	width := 1000
	keyRange := 10

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := 0
	started := false
	wg := sync.WaitGroup{}

	for i := 0; i < width; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			ready++
			if ready == width {
				started = true
				cond.Broadcast()
			} else {
				for !started {
					cond.Wait()
				}
			}
			mu.Unlock()

			key := fmt.Sprint(rand.Intn(keyRange) + 1)

			store.Set(key, key)

			_, _ = store.Get(key)

			store.Delete(key)

			_, _ = store.Get(key)
		}()
	}

	wg.Wait()

	for i := 1; i <= keyRange; i++ {
		v, ok := store.Get(fmt.Sprint(i))
		if ok {
			t.Errorf("key %q: got %q, but deleted", i, v)
		}
	}
}
