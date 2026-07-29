package queue_impl

import (
	"fmt"
)

type Queue struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	val  int
	next *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int) {
	node := &Node{
		val: val,
	}

	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		prevTail := q.tail
		q.tail = node
		prevTail.next = q.tail
	}

	q.size++
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Dequeue() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("empty queue")
	}

	val := q.head.val
	q.head = q.head.next
	q.size--
	return val, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("empty queue")
	}

	return q.head.val, nil
}
