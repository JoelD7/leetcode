package insert_delete_get_random_o1

import (
	"math/rand"
)

type RandomizedSet struct {
	m    map[int]int
	keys []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    make(map[int]int),
		keys: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]
	if ok {
		return false
	}

	this.m[val] = val
	this.updateKeys()
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.m[val]
	if ok {
		delete(this.m, val)
		this.updateKeys()
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	maxIndex := len(this.keys)

	if len(this.keys) == 1 {
		return this.m[this.keys[0]]
	}

	index, val := 0, 0
	ok := false

	for !ok {
		index = rand.Intn(maxIndex)
		val, ok = this.m[this.keys[index]]
		if ok {
			return val
		}
	}

	return 0
}

func (this *RandomizedSet) updateKeys() {
	this.keys = make([]int, 0)
	for k, _ := range this.m {
		this.keys = append(this.keys, k)
	}
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
