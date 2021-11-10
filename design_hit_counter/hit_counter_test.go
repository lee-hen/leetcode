package design_hit_counter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	obj := Constructor()
	obj.Hit(1)
	obj.Hit(2)
	obj.Hit(3)
	require.Equal(t, 3, obj.GetHits(4))

	obj.Hit(300)
	require.Equal(t, 4, obj.GetHits(300))
	require.Equal(t, 3, obj.GetHits(301))
}

func TestCase2(t *testing.T) {
	obj := Constructor()
	obj.Hit(1)
	obj.Hit(1)
	obj.Hit(1)
	obj.Hit(300)
	require.Equal(t, 4, obj.GetHits(300))
	obj.Hit(300)
	require.Equal(t, 5, obj.GetHits(300))
	obj.Hit(301)
	require.Equal(t, 3, obj.GetHits(301))
}

func TestCase3(t *testing.T) {
	obj := Constructor()
	obj.Hit(100)
	obj.Hit(282)
	obj.Hit(411)
	obj.Hit(609)
	obj.Hit(620)
	obj.Hit(744)
	require.Equal(t, 3, obj.GetHits(879))
	obj.Hit(956)
	require.Equal(t, 2, obj.GetHits(976))
	obj.Hit(998)
	require.Equal(t, 2, obj.GetHits(1055))
}
