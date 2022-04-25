package merge_intervals

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	require.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, Merge(intervals))

	intervals = [][]int{{1, 4}, {4, 5}}
	require.Equal(t, [][]int{{1, 5}}, Merge(intervals))
}
