package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	out := make([]string, 0)

	if root == nil {
		return "nil"
	}

	var node *TreeNode

	queue := NewQueue()
	queue.Enqueue(root)

	for queue.size > 0 {
		nodesInLevel := queue.size

		for i := 0; i < nodesInLevel && queue.size > 0; i++ {
			node = queue.Dequeue()
			if node == nil {
				out = append(out, "nil")
				continue
			}

			out = append(out, strconv.Itoa(node.Val))

			if node.Left != nil {
				queue.Enqueue(node.Left)
			} else {
				queue.Enqueue(nil)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			} else {
				queue.Enqueue(nil)
			}
		}
	}

	return strings.Join(out, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	inputArr := strings.Split(data, ",")

	//1,2,3,nil,nil,4,5,6,7,nil,nil,nil,nil,nil,nil
	if inputArr[0] == "nil" {
		return nil
	}

	val, _ := strconv.Atoi(inputArr[0])
	root := &TreeNode{
		Val: val,
	}

	queue := NewQueue()
	queue.Enqueue(root)
	i := 1

	for queue.size > 0 && i < len(inputArr) {
		cur := queue.Dequeue()

		if inputArr[i] != "nil" {
			val, _ = strconv.Atoi(inputArr[i])
			cur.Left = &TreeNode{Val: val}
			queue.Enqueue(cur.Left)
		}
		i++

		if inputArr[i] != "nil" && i < len(inputArr) {
			val, _ = strconv.Atoi(inputArr[i])
			cur.Right = &TreeNode{Val: val}
			queue.Enqueue(cur.Right)
		}
		i++
	}

	return root
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  *TreeNode
	next *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val *TreeNode) {
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

func (q *Queue) Dequeue() *TreeNode {
	if q.size == 0 {
		return nil
	}
	val := q.head.val
	q.head = q.head.next
	q.size--

	return val
}
