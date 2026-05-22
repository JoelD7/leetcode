package main

type MaxHeap struct {
	arr []int
}

func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap(stones)
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	for len(heap.arr) > 1 {
		x := heap.remove()
		y := heap.remove()
		res := abs(x, y)
		if res > 0 {
			heap.insert(res)
		}
	}

	if len(heap.arr) == 0 {
		return 0
	}

	return heap.arr[0]
}

func NewMaxHeap(vals []int) *MaxHeap {
	heap := &MaxHeap{}
	for _, val := range vals {
		heap.insert(val)
	}

	return heap
}

func (h *MaxHeap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MaxHeap) remove() int {
	val := h.arr[0]

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	h.heapifyDown(0)

	return val
}

func (h *MaxHeap) heapifyDown(index int) {
	left, right := leftChild(index), rightChild(index)
	largestChild := 0

	for left < len(h.arr) {
		if left == len(h.arr)-1 || h.arr[left] > h.arr[right] {
			largestChild = left
		} else {
			largestChild = right
		}

		if h.arr[largestChild] > h.arr[index] {
			h.swap(largestChild, index)
			index = largestChild
			left, right = leftChild(index), rightChild(index)
		} else {
			break
		}
	}
}

func (h *MaxHeap) heapifyUp(index int) {
	parentIdx := parent(index)

	for h.arr[parentIdx] < h.arr[index] {
		h.swap(parentIdx, index)
		index = parentIdx
		parentIdx = parent(index)
	}
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
