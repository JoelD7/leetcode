package design_circular_queue

type MyCircularQueue struct {
	//capacity of the queue. Indicates the max amount of elements it can hold
	cap int

	//length of the queue. Indicates the amount of elements the queue currently has
	len int

	head int
	tail int
	arr  []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap:  k,
		arr:  make([]int, k),
		tail: -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.tail == this.cap-1 {
		this.tail = 0
	} else {
		this.tail++
	}

	this.arr[this.tail] = value
	this.len++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	//This means that the "enqueueing" has "circled back"
	if this.head == this.cap-1 {
		this.head = 0
	} else {
		this.head++
	}

	this.len--

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
