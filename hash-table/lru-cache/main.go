package lru_cache

import (
	"math"
)

type LRUCache struct {
	incr       int
	lowestIncr int
	cache      map[int]int
	keyByIncr  map[int]int
	incrByKey  map[int]int
	cap        int
	length     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		incr:       0,
		lowestIncr: math.MaxInt32,
		length:     0,
		cap:        capacity,
		cache:      make(map[int]int),
		keyByIncr:  make(map[int]int),
		incrByKey:  make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if !ok {
		return -1
	}

	var keyIncr int
	keyIncr, ok = this.incrByKey[key]
	if ok {
		if keyIncr < this.lowestIncr {
			this.lowestIncr = keyIncr + 1
		}
		delete(this.keyByIncr, keyIncr)
	}

	this.incr++
	this.keyByIncr[this.incr] = key
	this.incrByKey[key] = this.incr
	return val
}

func (this *LRUCache) Put(key int, value int) {
	var ok bool
	_, ok = this.cache[key]
	if ok {
		this.cache[key] = value
		this.incr++

		cur, _ := this.incrByKey[key]
		delete(this.keyByIncr, cur)
		this.keyByIncr[this.incr] = key

		this.incrByKey[key] = this.incr
		return
	}

	if this.length == this.cap {
		ok = false
		var k int
		for !ok {
			k, ok = this.keyByIncr[this.lowestIncr]
			this.lowestIncr++
		}
		delete(this.cache, k)
		delete(this.incrByKey, k)
	}

	this.cache[key] = value
	this.incr++
	this.keyByIncr[this.incr] = key
	if this.incr < this.lowestIncr {
		this.lowestIncr = this.incr
	}
	this.incrByKey[key] = this.incr
	if this.length < this.cap {
		this.length++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
