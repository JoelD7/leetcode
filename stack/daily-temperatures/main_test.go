package daily_temperatures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	c := require.New(t)

	t.Run("temperatures = [73,74,75,71,69,72,76,73]", func(t *testing.T) {
		c.Equal([]int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	})

	t.Run("temperatures = [30, 40, 50, 60]", func(t *testing.T) {
		c.Equal([]int{1, 1, 1, 0}, dailyTemperatures([]int{30, 40, 50, 60}))
	})

	t.Run("temperatures = [30, 60, 90]", func(t *testing.T) {
		c.Equal([]int{1, 1, 0}, dailyTemperatures([]int{30, 60, 90}))
	})

	t.Run("temperatures = [a lot of 34s and 47s]", func(t *testing.T) {
		c.Equal([]int{2, 1, 0, 0, 3, 2, 1, 0, 1, 0, 0, 0, 3, 2, 1, 0, 0, 1, 0, 0, 4, 3, 2, 1, 0, 2, 1, 0, 3, 2, 1, 0, 0, 0}, dailyTemperatures([]int{34, 34, 47, 47, 34, 34, 34, 47, 34, 47, 47, 47, 34, 34, 34, 47, 47, 34, 47, 47, 34, 34, 34, 34, 47, 34, 34, 47, 34, 34, 34, 47, 47, 47}))
	})

}
