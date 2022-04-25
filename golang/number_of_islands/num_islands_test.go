package number_of_islands

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	require.Equal(t, 3, NumIslands(grid))
}
