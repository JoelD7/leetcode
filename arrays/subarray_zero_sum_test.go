package arrays

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindSubarrayZeroSum(t *testing.T) {
	c := require.New(t)

	fmt.Println(getPrefixSum([]int{4, 2, -3, 1, 6}))     //0,3
	fmt.Println(getPrefixSum([]int{5, 1, 2, -1, -1, 6})) //1,4
	fmt.Println(getPrefixSum([]int{8, 9, 3, 1, -6, 6}))  //3,5

	c.True(findSubarrayZeroSum([]int{4, 2, -3, 1, 6}))     //1,3
	c.True(findSubarrayZeroSum([]int{5, 1, 2, -1, -1, 6})) //2,4
	c.True(findSubarrayZeroSum([]int{8, 9, 3, 1, -6, 6}))  //4,5
	c.True(findSubarrayZeroSum([]int{4, 2, 0, 1, 6}))

}
