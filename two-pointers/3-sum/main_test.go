package three_sum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {

	fmt.Printf("Expected: [[-1,-1,2],[-1,0,1]], Got: %+v\n", threeSumRefresh([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("Expected: [[0,0,0]], Got: %+v\n", threeSumRefresh([]int{0, 0, 0, 0}))
	fmt.Printf("Expected: [[-1,0,1]], Got: %+v\n", threeSumRefresh([]int{1, -1, -1, 0}))
}
