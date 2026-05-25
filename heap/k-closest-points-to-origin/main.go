package main

import (
	"math"
)

type MaxHeap struct {
	arr []Point
}

type Point struct {
	distance   float64
	coordinate []int
}

func (h *MaxHeap) insert(val Point) {
	h.arr = append(h.arr, val)
	h.heapifyUp(h.size() - 1)
}

func (h *MaxHeap) remove() Point {
	val := h.arr[0]
	h.arr[0] = h.arr[h.size()-1]
	h.arr = h.arr[:h.size()-1]
	h.heapifyDown(0)
	return val
}

func (h *MaxHeap) size() int {
	return len(h.arr)
}

func (h *MaxHeap) heapifyUp(index int) {
	parentIdx := parent(index)
	for h.arr[parentIdx].distance < h.arr[index].distance {
		h.swap(parentIdx, index)
		index = parentIdx
		parentIdx = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	left, right := leftChild(index), rightChild(index)
	largestChild := 0

	for left < h.size() {
		if left == h.size()-1 || h.arr[left].distance > h.arr[right].distance {
			largestChild = left
		} else {
			largestChild = right
		}

		if h.arr[largestChild].distance > h.arr[index].distance {
			h.swap(largestChild, index)
			index = largestChild

			left, right = leftChild(index), rightChild(index)
		} else {
			break
		}
	}
}

func (h *MaxHeap) peek() Point {
	return h.arr[0]
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func distanceToOrigin(x, y int) float64 {
	return math.Sqrt(float64(x*x) + float64(y*y))
}

func kClosest(points [][]int, k int) [][]int {
	heap := &MaxHeap{}
	var point Point

	for i := 0; i < len(points); i++ {
		point = Point{
			distance:   distanceToOrigin(points[i][0], points[i][1]),
			coordinate: []int{points[i][0], points[i][1]},
		}

		if heap.size() < k || point.distance < heap.peek().distance {
			heap.insert(point)

			if heap.size() > k {
				heap.remove()
			}
		}
	}

	res := make([][]int, 0)
	for heap.size() > 0 {
		res = append(res, heap.remove().coordinate)
	}

	return res
}
