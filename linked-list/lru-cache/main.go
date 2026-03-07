package main

type LRUCache struct {
	cap   int
	size  int
	head  *Node
	tail  *Node
	table map[int]*Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{prev: nil}
	tail := &Node{next: nil}

	head.next = tail
	tail.prev = head

	return LRUCache{
		cap:   capacity,
		table: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.table[key]
	if !ok {
		return -1
	}

	this.moveToHead(node)

	return node.value
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	tail := this.tail.prev
	this.removeNode(tail)
	return tail
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.table[key]; exists {
		node.value = value
		this.moveToHead(node)
		return
	}

	newNode := &Node{
		key:   key,
		value: value,
	}

	this.table[key] = newNode
	this.addToHead(newNode)
	this.size++

	if this.size > this.cap {
		tail := this.removeTail()
		delete(this.table, tail.key)
		this.size--
	}
}
