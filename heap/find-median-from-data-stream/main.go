package main

import (
	"sort"
)

type MedianFinder struct {
	arr []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		arr: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.arr = append(this.arr, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.arr)

	length := len(this.arr)
	mid := length / 2

	if length%2 == 0 {
		return float64(this.arr[mid]+this.arr[(mid)-1]) / 2
	}

	return float64(this.arr[mid])
}
