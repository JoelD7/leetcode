package time_based_key_value_store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeMap(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		timeMap := Constructor()
		timeMap.Set("foo", "bar", 1)
		assert.Equal(t, "bar", timeMap.Get("foo", 1))
		assert.Equal(t, "bar", timeMap.Get("foo", 3))
		timeMap.Set("foo", "bar2", 4)
		assert.Equal(t, "bar2", timeMap.Get("foo", 4))
		assert.Equal(t, "bar2", timeMap.Get("foo", 5))
	})

	t.Run("Test case 2", func(t *testing.T) {
		timeMap := Constructor()
		timeMap.Set("love", "high", 10)
		timeMap.Set("love", "low", 20)
		assert.Equal(t, "", timeMap.Get("love", 5))
		assert.Equal(t, "high", timeMap.Get("love", 10))
		assert.Equal(t, "high", timeMap.Get("love", 15))
		assert.Equal(t, "low", timeMap.Get("love", 20))
		assert.Equal(t, "low", timeMap.Get("love", 25))
	})
}
