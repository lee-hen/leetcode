package lru_cache

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	lru list.List
	cache map[int]*list.Element
	deleted map[*list.Element]bool
}

func Constructor(capacity int) LRUCache {
	c := new(LRUCache)
	c.capacity = capacity
	c.cache = make(map[int]*list.Element)
	c.deleted = make(map[*list.Element]bool)
	return *c
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.cache[key]; ok {
		ele := c.cache[key]

		if c.deleted[ele] {
			delete(c.cache, key)
			delete(c.deleted, ele)
			return -1
		} else {
			c.lru.Remove(ele)
			c.lru.PushBack(ele.Value.(int))
			c.cache[key] = c.lru.Back()
			return ele.Value.(int)
		}
	}

	return -1
}

func (c *LRUCache) Put(key int, value int)  {
	if ele, ok := c.cache[key]; ok {
		c.lru.Remove(ele)
	}

	c.lru.PushBack(value)
	c.cache[key] = c.lru.Back()

	if c.lru.Len() > c.capacity  {
		ele := c.lru.Front()
		c.deleted[ele] = true
		c.lru.Remove(ele)
	}
}
