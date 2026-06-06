package main

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

type MinHeap struct {
	arr []int
}

type MaxHeap struct {
	arr []int
}

func newMinHeap() *MinHeap {
	return &MinHeap{
		arr: make([]int, 0),
	}
}

func (h *MinHeap) size() int {
	return len(h.arr)
}

func (h *MinHeap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(h.size() - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	parentIdx := parent(index)
	for h.arr[parentIdx] > h.arr[index] {
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

func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *MinHeap) peek() int {
	return h.arr[0]
}

func (h *MinHeap) remove() int {
	val := h.arr[0]
	h.arr[0] = h.arr[h.size()-1]
	h.arr = h.arr[:h.size()-1]
	h.heapifyDown(0)
	return val
}

func (h *MinHeap) heapifyDown(index int) {
	left, right := leftChild(index), rightChild(index)
	minChild := 0

	for left < h.size() {
		if left == h.size()-1 || h.arr[left] < h.arr[right] {
			minChild = left
		} else {
			minChild = right
		}

		if h.arr[minChild] < h.arr[index] {
			h.swap(minChild, index)
			index = minChild
			left, right = leftChild(index), rightChild(index)
		} else {
			break
		}
	}
}

func newMaxHeap() *MaxHeap {
	return &MaxHeap{
		arr: make([]int, 0),
	}
}

func (h *MaxHeap) insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	parentIdx := parent(index)

	for h.arr[parentIdx] < h.arr[index] {
		h.swap(parentIdx, index)
		index = parentIdx
		parentIdx = parent(index)
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *MaxHeap) Peek() int {
	return h.arr[0]
}

func (h *MaxHeap) Remove() int {
	val := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	h.heapifyDown(0)

	return val
}

func (h *MaxHeap) heapifyDown(index int) {
	left, right := leftChild(index), rightChild(index)
	maxChild := 0

	for left < h.size() {
		if left == h.size()-1 || h.arr[left] > h.arr[right] {
			maxChild = left
		} else {
			maxChild = right
		}

		if h.arr[maxChild] > h.arr[index] {
			h.swap(maxChild, index)
			index = maxChild
			left, right = leftChild(index), rightChild(index)
		} else {
			break
		}
	}
}

func (h *MaxHeap) size() int {
	return len(h.arr)
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: newMinHeap(),
		maxHeap: newMaxHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.maxHeap.insert(num)
	this.minHeap.insert(this.maxHeap.Remove())

	if this.minHeap.size() > this.maxHeap.size() {
		this.maxHeap.insert(this.minHeap.remove())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.size() > this.minHeap.size() {
		return float64(this.maxHeap.Peek())
	}

	return float64(this.maxHeap.Peek()+this.minHeap.peek()) / 2.0
}
