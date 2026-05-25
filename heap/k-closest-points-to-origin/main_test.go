package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// [[-5,4],[-6,-5],[4,6]]
func TestKClosest(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		assert.ElementsMatch(t, [][]int{{-5, 4}, {4, 6}}, kClosest([][]int{{-5, 4}, {-6, -5}, {4, 6}}, 2))
	})

	t.Run("Test case 2", func(t *testing.T) {
		//	[[-5,4],[-3,2],[0,1],[-3,7],[-2,0],[-4,-6],[0,-5]]
		input := [][]int{{-5, 4}, {-3, 2}, {0, 1}, {-3, 7}, {-2, 0}, {-4, -6}, {0, 5}}
		assert.ElementsMatch(t, [][]int{{-5, 4}, {-3, 2}, {0, 1}, {-2, 0}, {-4, -6}, {0, 5}}, kClosest(input, 6))
		//for i := 0; i < len(input); i++ {
		//	fmt.Printf("[%d, %d] = %f\n", input[i][0], input[i][1], distanceToOrigin(input[i][0], input[i][1]))
		//}
	})
}
