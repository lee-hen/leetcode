package count_of_smaller_numbers_after_self

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	nums := []int{5, 2, 6, 1}
	require.Equal(t, countSmaller(nums), CountSmaller(nums))

	// 0  1  2  3  4  5  6
	// 8, 5, 2, 9, 5, 6, 3

	//  0      1      2     3
	// [{5 1} {8 0} | {2 2} {7 3}
	nums = []int{8, 5, 2, 9, 5, 6, 3}
	require.Equal(t, countSmaller(nums), CountSmaller(nums))

	nums = []int{11, 9, 6, 7, 4, 5, 8, 20, 12}
	require.Equal(t, countSmaller(nums), CountSmaller(nums))

	nums = []int{5, 2, 6, 1, 4, 9, 3, 8, 7, 6, 4, 5, 1}
	require.Equal(t, countSmaller(nums), CountSmaller(nums))
}
