package longest_increasing_path

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	require.Equal(t, longestIncreasingPath([][]int{{6, 4, 5}, {3, 2, 6}, {2, 2, 1}}), LongestIncreasingPath([][]int{{6, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
}
