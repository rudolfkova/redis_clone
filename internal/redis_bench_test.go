package redis

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	keyRange = 10000
)

func BenchmarkSetGet(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		key := fmt.Sprint(r.Intn(keyRange))
		for pb.Next() {
			store.Set(key, key)
			_, _ = store.Get(key)
		}
	})
}

func BenchmarkSet(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for pb.Next() {
			key := fmt.Sprint(r.Intn(keyRange))
			store.Set(key, key)
		}
	})
}

func BenchmarkGet(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		key := fmt.Sprint(r.Intn(keyRange))
		for pb.Next() {
			_, _ = store.Get(key)
		}
	})
}

func BenchmarkDel(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		key := fmt.Sprint(r.Intn(keyRange))
		for pb.Next() {
			store.Delete(key)
		}
	})
}

func BenchmarkSetGetDel(b *testing.B) {
	store := NewStorage()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		key := fmt.Sprint(r.Intn(keyRange))
		for pb.Next() {
			store.Set(key, key)
			_, _ = store.Get(key)
			store.Delete(key)
		}
	})
}
