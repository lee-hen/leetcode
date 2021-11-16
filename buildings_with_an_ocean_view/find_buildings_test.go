package buildings_with_an_ocean_view

import (
	"github.com/lee-hen/leecode/lib"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	heights := []int{4,2,3,1}
	require.Equal(t, []int{0,2,3}, FindBuildings(heights))
}

func TestCase2(t *testing.T) {
	heights := []int{4,3,2,1}
	require.Equal(t, []int{0,1,2,3}, FindBuildings(heights))
}

func TestCase3(t *testing.T) {
	heights := []int{1,3,2,4}
	require.Equal(t, []int{3}, FindBuildings(heights))
}

func TestCase4(t *testing.T) {
	heights := []int{2,2,2,2}
	require.Equal(t, []int{3}, FindBuildings(heights))
}

func TestFindBuildings(t *testing.T) {
	heights := lib.RandomArray(10, 0, 100)
	t.Log(heights)
	t.Log(FindBuildings(heights))
}
