package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	lru      *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		lru:      list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		this.lru.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok {
		ele.Value.(*entry).value = value
		this.lru.MoveToFront(ele)
		return
	}

	if this.lru.Len() == this.capacity {
		ele := this.lru.Back()
		delete(this.cache, ele.Value.(*entry).key)
		this.lru.Remove(ele)
	}

	ele := this.lru.PushFront(&entry{key, value})
	this.cache[key] = ele
}

func (this *LRUCache) Print() {
	for e := this.lru.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*entry).key, e.Value.(*entry).value)
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)

	cache.Get(3)

	cache.Print()
}
