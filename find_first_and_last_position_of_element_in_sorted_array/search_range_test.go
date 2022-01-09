package find_first_and_last_position_of_element_in_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	require.Equal(t, []int{3, 4}, SearchRange(nums, target))

	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	require.Equal(t, []int{-1, -1}, SearchRange(nums, target))

	nums = []int{}
	target = 0
	require.Equal(t, []int{-1, -1}, SearchRange(nums, target))
}
