package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	// Represents:
	//      1
	//     / \
	//    2   3
	//     \
	//      5
	root := BuildTree([]*int{
		Ptr(1), Ptr(2), Ptr(3), nil, Ptr(5),
	})

	root.Print()

	assert.NotNil(t, root)
}
