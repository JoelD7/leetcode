package lru_cache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// ["LRUCache","put","put","get","put","get","put","get","get","get"]
// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
func TestLRUCache(t *testing.T) {
	c := require.New(t)

	o := Constructor(2)
	o.Put(1, 1)
	o.Put(2, 2)
	c.Equal(1, o.Get(1))
	o.Put(3, 3)
	c.Equal(-1, o.Get(2))
	o.Put(4, 4)
	c.Equal(-1, o.Get(1))
	c.Equal(3, o.Get(3))
	c.Equal(4, o.Get(4))
}
