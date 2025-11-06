package combinations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombine(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		k        int
		expected [][]int
	}{
		{
			name: "Example 1: n=4, k=2",
			n:    4,
			k:    2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		{
			name: "Example 2: n=1, k=1",
			n:    1,
			k:    1,
			expected: [][]int{
				{1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := combine(tc.n, tc.k)

			require.ElementsMatch(t, tc.expected, actual)
		})
	}
}
