package number_of_distinct_islands

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumDistinctIslands(t *testing.T) {
	grid := [][]int{{1,1,0,0,0},{1,1,0,0,0},{0,0,0,1,1},{0,0,0,1,1}}
	require.Equal(t, 1, NumDistinctIslands(grid))
	require.Equal(t, 1, numDistinctIslands(grid))
}

func TestNumDistinctIslands1(t *testing.T) {
	grid := [][]int{{1,1,0,1,1},{1,0,0,0,0},{0,0,0,0,1},{1,1,0,1,1}}
	require.Equal(t, 3, NumDistinctIslands(grid))
	require.Equal(t, 3, numDistinctIslands(grid))
}
