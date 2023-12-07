package min_stack

/*
Notes

- Min
	- Min is pointer to the minimum value
	- on each push, update min if val < min
*/

type Node struct {
	prev *Node
	val  int
}

type MinStack struct {
	head     *Node
	minStack *MinStack
	min      int
	length   int
}

func Constructor() MinStack {
	return MinStack{
		head: nil,
		minStack: &MinStack{
			head:   nil,
			length: 0,
		},
		length: 0,
	}
}

func (this *MinStack) Push(val int) {
	this.length++

	if this.head == nil {
		newHead := &Node{nil, val}
		this.head = newHead
		this.min = val
		return
	}

	curHead := &Node{this.head.prev, this.head.val}
	this.head.val = val
	this.head.prev = curHead

	if val < this.min {
		this.minStack.Push(val)
	}
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}

	this.length--

	this.head = this.head.prev
}

func (this *MinStack) Top() int {
	return this.head.val
}

// GetMin returns the minimum value in the stack without removing it.
func (this *MinStack) GetMin() int {

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
