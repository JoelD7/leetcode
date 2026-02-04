package main

import (
	"math"
)

type Queue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func NewQueue() Queue {
	return Queue{}
}

func (this *Queue) isEmpty() bool {
	return this.size == 0
}

func (this *Queue) Enqueue(x int) {
	oldTail := this.tail
	this.tail = &Node{
		val: x,
	}

	if this.isEmpty() {
		this.head = this.tail
	} else {
		oldTail.next = this.tail
	}

	this.size++
}

func (this *Queue) Dequeue() int {
	if this.isEmpty() {
		return -1
	}

	val := this.head.val
	this.head = this.head.next
	this.size--

	return val
}

func (this *Queue) Peek() int {
	return this.head.val
}

func maxSlidingWindow(nums []int, k int) []int {
	results := make([]int, 0)
	max := math.MinInt32

	queue := NewQueue()

	for i := 0; i < k; i++ {
		queue.Enqueue(nums[i])

		if nums[i] > max {
			max = nums[i]
		}
	}
	results = append(results, max)

	for i := k; i < len(nums); i++ {
		queue.Dequeue()
		queue.Enqueue(nums[i])
		if nums[i] > max {
			max = nums[i]
		}
		results = append(results, max)
	}

	return results
}
