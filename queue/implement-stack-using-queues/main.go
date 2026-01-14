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

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int) {
	newNode := &Node{
		val: val,
	}

	if q.head == nil {
		q.head = newNode
		q.tail = q.head
	} else {
		oldTail := q.tail
		q.tail = newNode
		oldTail.next = q.tail
	}

	q.size++
}

func (q *Queue) Dequeue() int {
	if q.size == 0 {
		return -1
	}
	val := q.head.val
	q.head = q.head.next
	q.size--

	return val
}

func (q *Queue) Peek() int {
	return q.head.val
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

type MyStack struct {
	queue *Queue
	size  int
}

func Constructor() MyStack {
	return MyStack{
		queue: NewQueue(),
	}
}

func (this *MyStack) Push(x int) {
	this.queue.Enqueue(x)

	size := this.queue.size
	for i := 0; i < size-1; i++ {
		this.queue.Enqueue(this.queue.Dequeue())
	}

	this.size++
}

func (this *MyStack) Pop() int {
	this.size--
	return this.queue.Dequeue()
}

func (this *MyStack) Top() int {
	return this.queue.Peek()
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}
