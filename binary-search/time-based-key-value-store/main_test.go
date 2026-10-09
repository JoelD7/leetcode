package time_based_key_value_store

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	t.Run("Get before any set", func(t *testing.T) {
		timeMap := Constructor()
		assert.Equal(t, "", timeMap.Get("foo", 1))
	})

	t.Run("Multiple keys and intermediate timestamps", func(t *testing.T) {
		timeMap := Constructor()
		timeMap.Set("love", "high", 10)
		timeMap.Set("love", "low", 20)

		// Before first timestamp
		assert.Equal(t, "", timeMap.Get("love", 5))

		// Exact first timestamp
		assert.Equal(t, "high", timeMap.Get("love", 10))

		// Between timestamps
		assert.Equal(t, "high", timeMap.Get("love", 15))

		// Exact second timestamp
		assert.Equal(t, "low", timeMap.Get("love", 20))

		// After second timestamp
		assert.Equal(t, "low", timeMap.Get("love", 25))
	})

	t.Run("Different keys don't interfere with each other", func(t *testing.T) {
		timeMap := Constructor()
		timeMap.Set("key1", "val1", 10)
		timeMap.Set("key2", "val2", 20)

		assert.Equal(t, "val1", timeMap.Get("key1", 15))
		assert.Equal(t, "val2", timeMap.Get("key2", 25))
		assert.Equal(t, "", timeMap.Get("key1", 5))
		assert.Equal(t, "", timeMap.Get("key2", 15))
	})

	t.Run("Multiple updates to the same key over time", func(t *testing.T) {
		timeMap := Constructor()
		timeMap.Set("a", "b", 1)
		timeMap.Set("a", "c", 4)
		timeMap.Set("a", "d", 5)
		timeMap.Set("a", "e", 10)

		assert.Equal(t, "b", timeMap.Get("a", 2))
		assert.Equal(t, "c", timeMap.Get("a", 4))
		assert.Equal(t, "d", timeMap.Get("a", 7))
		assert.Equal(t, "e", timeMap.Get("a", 15))
	})

	t.Run("TEst case 2", func(t *testing.T) {
		timeMap := Constructor()
		timeMap.Set("foo", "bar", 5)
		timeMap.Set("foo", "bar2", 10)
		timeMap.Set("foo", "bar3", 14)
		assert.Equal(t, "", timeMap.Get("foo", 4))
		assert.Equal(t, "bar", timeMap.Get("foo", 6))
		assert.Equal(t, "bar2", timeMap.Get("foo", 11))
	})
}
