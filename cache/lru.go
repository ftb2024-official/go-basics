package cache

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Pair struct {
	key   int
	value int
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if v, exists := c.cache[key]; exists {
		c.list.MoveToFront(v)
		return v.Value.(*Pair).value
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if v, exists := c.cache[key]; exists {
		v.Value.(*Pair).value = value
		c.list.MoveToFront(v)
		return
	}

	if c.list.Len() >= c.capacity {
		// evict the least recently used element
		lru := c.list.Back()
		if lru != nil {
			delete(c.cache, lru.Value.(*Pair).key)
			c.list.Remove(lru)
		}
	}

	pair := &Pair{key: key, value: value}
	elem := c.list.PushFront(pair)
	c.cache[key] = elem
}

func Example1() {
	cache := New(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))

	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1

	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
