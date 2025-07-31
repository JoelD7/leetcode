package maximum_average_subarray

import (
	"github.com/stretchr/testify/require"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	c := require.New(t)

	t.Run("Very long input, n = 10000", func(t *testing.T) {
		nums, err := parseLargeInput()
		c.Nil(err)

		c.Equal(-25.14, roundUp(findMaxAverage(nums, 6514)))
	})

	t.Run("[0,1,1,3,3]", func(t *testing.T) {
		c.Equal(2.00, roundUp(findMaxAverage([]int{0, 1, 1, 3, 3}, 4)))
	})

	t.Run("[1,12,-5,-6,50,3]", func(t *testing.T) {
		c.Equal(12.75, roundUp(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)))
	})

	t.Run("[5]", func(t *testing.T) {
		c.Equal(5.00, roundUp(findMaxAverage([]int{5}, 1)))
	})
}

func parseLargeInput() ([]int, error) {
	data, err := os.ReadFile("large_input.txt")
	if err != nil {
		return nil, err
	}

	var num int
	nums := make([]int, 10_000)

	dataStr := strings.TrimSuffix(strings.TrimPrefix(string(data), "["), "]")
	itemsStr := strings.Split(dataStr, ",")
	for i, str := range itemsStr {
		num, err = strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		nums[i] = num
	}

	return nums, nil
}

func roundUp(val float64) float64 {
	return math.Ceil(val*100) / 100
}
