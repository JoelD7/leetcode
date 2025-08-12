package insert_delete_get_random_o1

import (
	"math/rand"
)

type RandomizedSet struct {
	m          map[int]int
	indexByVal map[int]int
	i          int
	keys       []int
	lastDelIdx int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:          make(map[int]int),
		indexByVal: make(map[int]int),
		keys:       make([]int, 0),
		lastDelIdx: -1,
		i:          0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.indexByVal[val]
	if ok {
		return false
	}

	if this.lastDelIdx >= 0 {
		this.m[this.lastDelIdx] = val
		this.indexByVal[val] = this.lastDelIdx
		this.lastDelIdx = -1
	} else {
		this.m[this.i] = val
		this.indexByVal[val] = this.i
		this.i++
	}

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.indexByVal[val]
	if ok {
		delete(this.m, i)
		this.lastDelIdx = i
		delete(this.indexByVal, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.i)
	val, ok := this.m[index]
	if ok {
		return val
	}

	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
