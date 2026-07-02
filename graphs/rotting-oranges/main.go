package main

type Queue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  []int
	next *Node
}

func NewQueue() Queue {
	return Queue{}
}

func (this *Queue) isEmpty() bool {
	return this.size == 0
}

func (this *Queue) Enqueue(x []int) {
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

func (this *Queue) Dequeue() []int {
	if this.isEmpty() {
		return []int{}
	}

	val := this.head.val
	this.head = this.head.next
	this.size--

	return val
}

func (this *Queue) Peek() []int {
	return this.head.val
}

func orangesRotting(grid [][]int) int {
	var freshCount int
	mins := -1
	dirs := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	queue := NewQueue()

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 2 {
				queue.Enqueue([]int{r, c})
			} else if grid[r][c] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	if queue.isEmpty() {
		return -1
	}

	for !queue.isEmpty() {
		size := queue.size

		for size > 0 {
			node := queue.Dequeue()
			size--

			r := node[0]
			c := node[1]

			for _, direction := range dirs {
				i := r + direction[0]
				j := c + direction[1]

				if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[i]) && grid[i][j] == 1 {
					grid[i][j] = 2
					queue.Enqueue([]int{i, j})
					freshCount--
				}
			}
		}

		mins++
	}

	if freshCount == 0 {
		return mins
	}

	return -1
}
