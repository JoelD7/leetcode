package main

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Insert(val int) {
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

func parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) Remove() int {
	val := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]

	h.heapifyDown(0)

	return val
}

func (h *MaxHeap) heapifyDown(index int) {
	left := leftChild(index)
	right := rightChild(index)
	largestChild := left

	if right < len(h.arr) && h.arr[left] < h.arr[right] {
		largestChild = right
	}

	for largestChild < len(h.arr) && h.arr[index] < h.arr[largestChild] {
		h.swap(index, largestChild)
		index = largestChild

		left = leftChild(index)
		right = rightChild(index)
		largestChild = left

		if left >= len(h.arr) {
			break
		}

		if right < len(h.arr) && h.arr[left] < h.arr[right] {
			largestChild = right
		}
	}
}

func (h *MaxHeap) isEmtpy() bool {
	return len(h.arr) == 0
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func leastInterval(tasks []byte, n int) int {
	heap := &MaxHeap{}
	var time int
	freq := make([]int, 26)

	for _, t := range tasks {
		freq[t-'A']++
	}

	for _, f := range freq {
		if f > 0 {
			heap.Insert(f)
		}
	}

	for !heap.isEmtpy() {
		cycle := n + 1
		store := make([]int, 0)
		taskCount := 0

		for cycle > 0 && !heap.isEmtpy() {
			taskFreq := heap.Remove()
			if taskFreq > 1 {
				store = append(store, taskFreq-1)
			}
			cycle--
			taskCount++
		}

		for i := 0; i < len(store); i++ {
			heap.Insert(store[i])
		}

		if heap.isEmtpy() {
			time += taskCount
		} else {
			time += n + 1
		}

	}

	return time
}
