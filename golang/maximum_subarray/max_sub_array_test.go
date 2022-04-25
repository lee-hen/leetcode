package maximum_subarray

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	require.Equal(t, maxSubArray(nums), MaxSubArray(nums))

	nums = []int{-1}
	require.Equal(t, maxSubArray(nums), MaxSubArray(nums))
}
