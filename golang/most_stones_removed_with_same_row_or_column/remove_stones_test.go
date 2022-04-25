package most_stones_removed_with_same_row_or_column

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_removeStones(t *testing.T) {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	require.Equal(t, removeStones1(stones), removeStones(stones))

	stones = [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}
	require.Equal(t, removeStones1(stones), removeStones(stones))

	stones = [][]int{{0, 0}}
	require.Equal(t, removeStones1(stones), removeStones(stones))

	stones = [][]int{{3, 2}, {3, 1}, {4, 4}, {1, 1}, {0, 2}, {4, 0}}
	require.Equal(t, removeStones1(stones), removeStones(stones))
}
