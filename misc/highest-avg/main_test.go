package highest_avg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetHighestGradedStudent(t *testing.T) {
	c := require.New(t)

	c.Equal(1, getHighestGradedStudent([][]int{{1, 95}, {1, 98}, {9, 92}, {9, 94}, {34, 89}, {34, 91}}))
	c.Equal(12, getHighestGradedStudent([][]int{{10, 78}, {12, 89}, {10, 90}, {12, 92}, {8, 88}}))
	c.Equal(4, getHighestGradedStudent([][]int{{5, 85}, {5, 90}, {4, 88}, {4, 92}, {4, 90}, {13, 75}, {13, 80}}))
	c.Equal(7, getHighestGradedStudent([][]int{{78, 85}, {78, 92}, {7, 90}, {7, 94}, {2, 88}, {2, 95}, {85, 82}}))
	c.Equal(103, getHighestGradedStudent([][]int{{1, 85}, {1, 90}, {103, 92}, {103, 95}, {3, 88}, {3, 90}, {44, 84}, {44, 87}}))
}
