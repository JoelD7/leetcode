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

	this.head.next = node
	this.head.next.prev = node
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
	newNode := &Node{
		key:   key,
		value: value,
	}

	if this.size < this.cap {
		this.size++
		this.table[key] = newNode
		this.addToHead(newNode)
		return
	}

	node := this.table[key]

	//New node
	if node == nil {
		tail := this.removeTail()
		delete(this.table, tail.key)
		this.addToHead(newNode)

		//Updating existing node
	} else {
		this.moveToHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
