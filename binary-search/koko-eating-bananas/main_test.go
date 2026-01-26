package koko_eating_bananas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	t.Run("[30,11,23,4,20] h = 6", func(t *testing.T) {
		assert.Equal(t, 23, minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	})

	t.Run("[312884470], h = 968709470", func(t *testing.T) {
		assert.Equal(t, 1, minEatingSpeed([]int{312884470}, 968709470))
	})
}
