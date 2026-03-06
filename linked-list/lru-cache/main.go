package main

type LRUCache struct {
	cap   int
	size  int
	head  *Node
	tail  *Node
	table map[int]int
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		table: make(map[int]int),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.table[key]
	if !ok {
		return -1
	}

	this.addNodeToList(key, val)

	return val
}

func (this *LRUCache) addNodeToList(key, val int) {
	if this.head == nil {
		this.head = &Node{
			key:   key,
			value: val,
		}
		this.tail = this.head
	} else {
		prevTail := this.tail
		newNode := &Node{
			key:   key,
			value: val,
			prev:  prevTail,
		}
		prevTail.next = newNode
		this.tail = newNode
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.size < this.cap {
		this.size++
		this.table[key] = value
		this.addNodeToList(key, value)
		return
	}

	count := 0
	cur := this.tail
	visited := make(map[int]bool)

	for count < this.size-1 {
		ok := visited[cur.key]
		if ok {
			//This shouldn't produce nil pointer errors
			cur.prev.next = cur.next
			cur = cur.prev
			continue
		}

		visited[cur.key] = true
		cur = cur.prev
		count++
	}

	this.head = cur.next

	this.table[key] = value
	this.addNodeToList(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
