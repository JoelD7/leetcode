package main

type Queue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

// For simplicity, nodes go from 0 to N, where N = len(adjList)
func bfsGraphs(adjList [][]int) ([]int, int) {
	queue := NewQueue()
	result := make([]int, 0)

	if len(adjList) == 0 {
		return result, 0
	}

	visited := make([]bool, len(adjList))

	queue.Enqueue(0)
	visited[0] = true
	lvl := 0

	for !queue.isEmpty() {
		size := queue.size

		for i := 0; i < size; i++ {
			node := queue.Dequeue()
			result = append(result, node)

			for _, neighbor := range adjList[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue.Enqueue(neighbor)
				}
			}
		}

		lvl++
	}

	return result, lvl
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
