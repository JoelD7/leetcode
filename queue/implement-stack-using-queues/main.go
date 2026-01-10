package implement_stack_using_queues

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
	val := this.head.val
	this.head = this.head.next
	this.size--
	return val
}

func (this *Queue) Peek() int {
	return this.head.val
}

type MyStack struct {
	queue        Queue
	reverseQueue Queue
	size         int
}

func Constructor() MyStack {
	return MyStack{
		queue:        NewQueue(),
		reverseQueue: NewQueue(),
	}
}

func (this *MyStack) Push(x int) {
	this.queue.Enqueue(x)

	elements := make([]int, 0)
	for !this.queue.isEmpty() {
		elements = append(elements, this.queue.Dequeue())
	}

	this.reverseQueue = NewQueue()

	i := len(elements) - 1
	j := 0
	for i >= 0 && j < len(elements) {
		this.reverseQueue.Enqueue(elements[i])
		this.queue.Enqueue(elements[j])
		i--
		j++
	}
	this.size++
}

func (this *MyStack) Pop() int {
	this.queue = NewQueue()
	reversedElements := make([]int, 0)

	val := this.reverseQueue.Dequeue()

	for !this.reverseQueue.isEmpty() {
		reversedElements = append(reversedElements, this.reverseQueue.Dequeue())
	}

	i := len(reversedElements) - 1
	j := 0
	for i >= 0 && j < len(reversedElements) {
		this.queue.Enqueue(reversedElements[i])
		this.reverseQueue.Enqueue(reversedElements[j])
		i--
		j++
	}

	this.size--

	return val
}

func (this *MyStack) Top() int {
	return this.reverseQueue.Peek()
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
