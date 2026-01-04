package redis

import (
	"hash/fnv"
	"sync"
)

const (
	shardsNum = 64
)

type Shard struct {
	storage map[string]string
	mu      sync.RWMutex
}

type Redis struct {
	shards []*Shard
}

func NewStorage() *Redis {
	r := &Redis{
		shards: make([]*Shard, shardsNum),
	}

	for s := range r.shards {
		r.shards[s] = &Shard{}
		r.shards[s].makeMap()
	}

	return r
}

func (s *Shard) makeMap() {
	s.storage = make(map[string]string)
}

func (r *Redis) Set(key string, value string) {
	hash := hashKey(key) % shardsNum

	r.shards[hash].mu.Lock()
	r.shards[hash].storage[key] = value
	r.shards[hash].mu.Unlock()

}

func (r *Redis) Get(key string) (string, bool) {
	hash := hashKey(key) % shardsNum

	r.shards[hash].mu.RLock()
	v, ok := r.shards[hash].storage[key]
	r.shards[hash].mu.RUnlock()

	return v, ok
}

func (r *Redis) Delete(key string) {
	hash := hashKey(key) % shardsNum

	r.shards[hash].mu.Lock()
	delete(r.shards[hash].storage, key)
	r.shards[hash].mu.Unlock()
}

func hashKey(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}
