package main

type KthLargest struct {
	arr []int
	k   int
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := KthLargest{
		k:   k,
		arr: make([]int, 0),
	}

	for i := 0; i < len(nums); i++ {
		kthLargest.Add(nums[i])
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if len(this.arr) < this.k || this.Peek() < val {
		this.arr = append(this.arr, val)
		this.heapifyUp(len(this.arr) - 1)

		if len(this.arr) > this.k {
			this.Remove()
		}
	}

	return this.Peek()
}

func (this *KthLargest) Remove() {
	this.arr[0] = this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]

	this.heapifyDown(0)
}

func (this *KthLargest) Peek() int {
	return this.arr[0]
}

func (this *KthLargest) heapifyUp(index int) {
	parentIdx := parent(index)

	for this.arr[parentIdx] > this.arr[index] {
		this.swap(index, parentIdx)
		index = parentIdx
		parentIdx = parent(index)
	}
}

func (this *KthLargest) heapifyDown(index int) {
	left := leftChild(index)
	right := rightChild(index)
	smallestChild := 0

	for left < len(this.arr) {
		if left == len(this.arr)-1 || this.arr[left] < this.arr[right] {
			smallestChild = left
		} else {
			smallestChild = right
		}

		if this.arr[smallestChild] < this.arr[index] {
			this.swap(smallestChild, index)
			index = smallestChild

			left = leftChild(index)
			right = rightChild(index)
		} else {
			break
		}
	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (this *KthLargest) swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}
