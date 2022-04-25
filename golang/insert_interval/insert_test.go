package insert_interval

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsert(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	require.Equal(t, [][]int{{1, 5}, {6, 9}}, Insert(intervals, newInterval))

	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 8}
	require.Equal(t, [][]int{{1, 2}, {3, 10}, {12, 16}}, Insert(intervals, newInterval))
}

func Test_insert(t *testing.T) {
	intervals := [][]int{{1, 5}}
	newInterval := []int{2, 3}
	require.Equal(t, [][]int{{1, 5}}, insert(intervals, newInterval))

	intervals = [][]int{{1, 5}}
	newInterval = []int{6, 8}
	require.Equal(t, [][]int{{1, 5}, {6, 8}}, insert(intervals, newInterval))

	intervals = [][]int{{1, 5}}
	newInterval = []int{0, 3}
	require.Equal(t, [][]int{{0, 5}}, insert(intervals, newInterval))

	intervals = [][]int{}
	newInterval = []int{5, 7}
	require.Equal(t, [][]int{{5, 7}}, insert(intervals, newInterval))
}
