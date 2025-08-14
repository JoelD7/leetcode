package insert_delete_get_random_o1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJoel(t *testing.T) {
	c := require.New(t)

	t.Run("case 1", func(t *testing.T) {
		o := Constructor()
		c.True(o.Insert(1))
		c.False(o.Remove(2))
		c.True(o.Insert(2))
		c.Equal(2, o.GetRandom())
		c.True(o.Remove(1))
		c.False(o.Insert(2))
		c.Equal(2, o.GetRandom())

	})

}
