package main

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	parentIdx := parent(index)
	for h.arr[index] < h.arr[parentIdx] {
		h.swap(index, parentIdx)
		index = parentIdx
		parentIdx = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
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

func (h *MinHeap) Peek() int { return h.arr[0] }

func (h *MinHeap) Remove() int {
	val := h.arr[0]

	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyDown(0)

	return val
}

func (h *MinHeap) heapifyDown(index int) {
	l := leftChild(index)
	r := rightChild(index)

	loChildIdx := l
	if r < len(h.arr) && h.arr[r] < h.arr[l] {
		loChildIdx = r
	}

	for loChildIdx < len(h.arr) && h.arr[index] > h.arr[loChildIdx] {
		h.swap(index, loChildIdx)
		index = loChildIdx

		l = leftChild(index)
		r = rightChild(index)

		loChildIdx = l
		if r < len(h.arr) && h.arr[r] < h.arr[l] {
			loChildIdx = r
		}
	}
}
