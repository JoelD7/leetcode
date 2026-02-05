package main

type Deque struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  Item
	next *Node
	prev *Node
}

type Item struct {
	val   int
	index int
}

func NewQueue() Deque {
	return Deque{}
}

func (this *Deque) isEmpty() bool {
	return this.size == 0
}

func (this *Deque) PushBack(x Item) {
	oldTail := this.tail
	this.tail = &Node{
		val:  x,
		prev: oldTail,
	}

	if this.isEmpty() {
		this.head = this.tail
	} else {
		oldTail.next = this.tail
	}

	this.size++
}

func (this *Deque) PopFront() Item {
	if this.isEmpty() {
		return Item{}
	}

	val := this.head.val
	this.head = this.head.next
	if this.head != nil {
		this.head.prev = nil
	}

	this.size--

	return val
}

func (this *Deque) PopBack() Item {
	if this.isEmpty() {
		return Item{}
	}

	val := this.tail.val
	this.tail = this.tail.prev
	if this.tail != nil {
		this.tail.next = nil
	}
	this.size--

	return val
}

func (this *Deque) Front() Item {
	return this.head.val
}

func (this *Deque) Back() Item {
	return this.tail.val
}

func maxSlidingWindow(nums []int, k int) []int {
	results := make([]int, 0)
	queue := NewQueue()
	var l, r int
	var cur Item

	for i := 0; i <= len(nums)-k; i++ {
		l = i
		r = i + k
		for j := l; j < r; j++ {
			cur = Item{
				val:   nums[j],
				index: j,
			}

			if queue.isEmpty() {
				queue.PushBack(cur)
				continue
			}

			for !queue.isEmpty() && queue.Back().val < cur.val {
				queue.PopBack()
			}

			queue.PushBack(cur)
		}

		for queue.Front().index < l {
			queue.PopFront()
		}

		results = append(results, queue.PopFront().val)
	}

	return results
}
