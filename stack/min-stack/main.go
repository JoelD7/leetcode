package min_stack

type MinStack struct {
	size int
	head Node
}

type Node struct {
	val  int
	min  int
	next *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.size == 0 {
		this.size++
		this.head = Node{
			val: val,
			min: val,
		}

		return
	}

	this.size++
	curHead := this.head
	this.head = Node{
		val:  val,
		next: &curHead,
		min:  getMin(curHead.min, val),
	}
}

func (this *MinStack) Pop() {
	this.size--

	var nextHead Node
	if this.head.next != nil {
		nextHead = *this.head.next
	}

	this.head = nextHead
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
