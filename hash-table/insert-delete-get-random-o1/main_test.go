package insert_delete_get_random_o1

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

//["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]
//[[],[1],[2],[2],[],[1],[2],[]]

func TestJoel(t *testing.T) {
	c := require.New(t)

	o := Constructor()
	c.True(o.Insert(1))
	c.True(o.Insert(2))
	c.True(o.Insert(3))
	c.True(o.Remove(1))
	c.True(o.Insert(1))
	c.True(o.Remove(1))
	c.True(o.Insert(1))
	fmt.Println(o.GetRandom())
}
