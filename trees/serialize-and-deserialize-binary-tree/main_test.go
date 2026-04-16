package main

import (
	"testing"

	"github.com/JoelD7/leetcode/trees/utils"
	"github.com/stretchr/testify/assert"
)

var ptr = utils.Ptr

func TestJoel(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		root := BuildTree([]*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), ptr(5)})

		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		deserialized := deser.deserialize(data)
		assert.NotNil(t, deserialized)
	})

	t.Run("Test case 2", func(t *testing.T) {
		root := BuildTree([]*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), ptr(5), ptr(6), ptr(7)})

		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		deserialized := deser.deserialize(data)
		assert.NotNil(t, deserialized)
	})

	t.Run("Example 2", func(t *testing.T) {
		root := BuildTree(nil)

		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		deserialized := deser.deserialize(data)
		assert.Nil(t, deserialized)
	})

	t.Run("Test case 3", func(t *testing.T) {

		root := BuildTree([]*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), ptr(5), ptr(6), ptr(7)})

		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		deserialized := deser.deserialize(data)
		assert.NotNil(t, deserialized)
	})
}

func BuildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			current.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			current.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}
