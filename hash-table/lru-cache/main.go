package lru_cache

type LRUCache struct {
	ageCount int
	oldest   int
	cache    map[int]int
	ageToKey map[int]int
	cap      int
	length   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		ageCount: 0,
		length:   0,
		oldest:   0,
		cap:      capacity,
		cache:    make(map[int]int),
		ageToKey: make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if !ok {
		return -1
	}

	return val
}

func (this *LRUCache) Put(key int, value int) {
	if this.length == this.cap {
		delete(this.cache, this.oldest)
		this.oldest++
		this.length = 0
	}

	this.cache[key] = value
	this.ageToKey[this.ageCount] = key
	this.ageCount++
	this.length++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
