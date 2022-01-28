package insert_delete_get_random

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	r := Constructor()
	r.Insert(0)
	r.Insert(1)
	r.Remove(0)
	r.Insert(2)
	r.Remove(1)

	require.Equal(t, 2, r.GetRandom())
}
