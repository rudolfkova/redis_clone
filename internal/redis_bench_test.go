package redis

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

const (
	keyRange = 10
)

var seedCounter uint64

func BenchmarkSetGet(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		seed := atomic.AddUint64(&seedCounter, 1)
		r := rand.New(rand.NewSource(int64(seed)))

		for pb.Next() {
			key := fmt.Sprint(r.Intn(keyRange))
			store.Set(key, key)
			_, _ = store.Get(key)
		}
	})
}

func BenchmarkSet(b *testing.B) {
	store := NewStorage()

	b.RunParallel(func(pb *testing.PB) {
		seed := atomic.AddUint64(&seedCounter, 1)
		r := rand.New(rand.NewSource(int64(seed)))
		for pb.Next() {
			key := fmt.Sprint(r.Intn(keyRange))
			store.Set(key, key)
		}
	})
}

func BenchmarkGet(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		seed := atomic.AddUint64(&seedCounter, 1)
		r := rand.New(rand.NewSource(int64(seed)))

		for pb.Next() {
			key := fmt.Sprint(r.Intn(keyRange))
			_, _ = store.Get(key)
		}
	})
}

func BenchmarkDel(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		seed := atomic.AddUint64(&seedCounter, 1)
		r := rand.New(rand.NewSource(int64(seed)))

		for pb.Next() {
			key := fmt.Sprint(r.Intn(keyRange))
			store.Delete(key)
		}
	})
}

func BenchmarkSetGetDel(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		seed := atomic.AddUint64(&seedCounter, 1)
		r := rand.New(rand.NewSource(int64(seed)))

		for pb.Next() {
			key := fmt.Sprint(r.Intn(keyRange))
			store.Set(key, key)
			_, _ = store.Get(key)
			store.Delete(key)
		}
	})
}
