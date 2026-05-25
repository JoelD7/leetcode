package main

type MinHeap struct {
	arr []int
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

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}

	for i := 0; i < len(nums); i++ {
		if h.size() < k || nums[i] > h.peek() {
			h.insert(nums[i])

			if h.size() > k {
				h.remove()
			}
		}
	}

	return h.peek()
}
