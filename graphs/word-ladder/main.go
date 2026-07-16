package main

type Queue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  string
	next *Node
}

func NewQueue() Queue {
	return Queue{}
}

func (this *Queue) isEmpty() bool {
	return this.size == 0
}

func (this *Queue) Enqueue(x string) {
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

func (this *Queue) Dequeue() string {
	if this.isEmpty() {
		return ""
	}

	val := this.head.val
	this.head = this.head.next
	this.size--

	return val
}

func (this *Queue) Peek() string {
	return this.head.val
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, s := range wordList {
		wordSet[s] = true
	}

	_, ok := wordSet[endWord]
	if !ok {
		return 0
	}

	adjList := make(map[string][]string)
	wordList = append(wordList, beginWord)
	var pattern string

	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern = word[:i] + "*" + word[i+1:]
			adjList[pattern] = append(adjList[pattern], word)
		}
	}

	queue := NewQueue()
	queue.Enqueue(beginWord)
	visited := make(map[string]bool)
	visited[beginWord] = true
	lvl := 1

	for !queue.isEmpty() {
		size := queue.size

		for i := 0; i < size; i++ {
			curWord := queue.Dequeue()

			if curWord == endWord {
				return lvl
			}

			//All nodes that share this pattern are adjacent to curWord
			for j := 0; j < len(curWord); j++ {
				pattern = curWord[:j] + "*" + curWord[j+1:]

				for _, neighbor := range adjList[pattern] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue.Enqueue(neighbor)
					}
				}
			}
		}

		lvl++
	}

	return 0
}
