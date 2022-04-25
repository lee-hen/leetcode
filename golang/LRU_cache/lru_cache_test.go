package lru_cache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLRUCache1(t *testing.T) {
	cacheSize := 2
	cache := Constructor(cacheSize)

	cache.Put(1, 1) // cache is {1=1}
	cache.Put(2, 2)  // cache is {1=1, 2=2}
	require.Equal(t, 1, cache.Get(1))
	cache.Put(3, 3)
	require.Equal(t, -1, cache.Get(2))  // return -1 (not found)
	cache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}

	require.Equal(t, -1, cache.Get(1))
	require.Equal(t, 3, cache.Get(3))
	require.Equal(t, 4, cache.Get(4))
}

func TestLRUCache2(t *testing.T) {
	cacheSize := 2
	cache := Constructor(cacheSize)

	cache.Put(2, 1)
	cache.Put(1, 1)
	require.Equal(t, 1, cache.Get(2))

	cache.Put(4, 1)
	require.Equal(t, -1, cache.Get(1))
	require.Equal(t, 1, cache.Get(2))
}

func TestLRUCache3(t *testing.T) {
	cacheSize := 2
	cache := Constructor(cacheSize)
	require.Equal(t, -1, cache.Get(2))

	cache.Put(2, 6)
	require.Equal(t, -1, cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)
	require.Equal(t, 2, cache.Get(1))
	require.Equal(t, 6, cache.Get(2))
}
