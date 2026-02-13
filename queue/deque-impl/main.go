package main

type DoublyEndedQueue interface {
	IsEmpty() bool
	PushFront(x int)
	PushBack(x int)
	PopFront() int
	PopBack() int
	Front() int
	Back() int
}

type Deque struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

func NewDeque() Deque {
	return Deque{}
}

func (this *Deque) IsEmpty() bool {
	return this.size == 0
}

func (this *Deque) PushFront(x int) {
	oldHead := this.head

	this.head = &Node{
		val:  x,
		next: oldHead,
	}

	if oldHead != nil {
		oldHead.prev = this.head
	} else {
		this.tail = this.head
	}

	this.size++
}

func (this *Deque) PushBack(x int) {
	oldTail := this.tail
	this.tail = &Node{
		val:  x,
		prev: oldTail,
	}

	if this.IsEmpty() {
		this.head = this.tail
	} else {
		oldTail.next = this.tail
	}

	this.size++
}

func (this *Deque) PopFront() int {
	if this.IsEmpty() {
		return -1
	}

	val := this.head.val
	this.head = this.head.next
	if this.head != nil {
		this.head.prev = nil
	}

	this.size--

	return val
}

func (this *Deque) PopBack() int {
	if this.IsEmpty() {
		return -1
	}

	val := this.tail.val
	this.tail = this.tail.prev
	if this.tail != nil {
		this.tail.next = nil
	}
	this.size--

	return val
}

func (this *Deque) Front() int {
	return this.head.val
}

func (this *Deque) Back() int {
	return this.tail.val
}
