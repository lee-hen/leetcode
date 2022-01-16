package max_product

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	require.Equal(t, 6, MaxProduct(nums))
	require.Equal(t, 6, maxProduct(nums))

	nums = []int{-2, 0, -1}
	require.Equal(t, 0, MaxProduct(nums))
	require.Equal(t, 0, maxProduct(nums))

	nums = []int{-2, 3, -4}
	require.Equal(t, 24, MaxProduct(nums))
	require.Equal(t, 24, maxProduct(nums))

	nums = []int{-3, -1, -1}
	require.Equal(t, 3, MaxProduct(nums))
	require.Equal(t, 3, maxProduct(nums))

	nums = []int{-3, 0, 1, -2}
	require.Equal(t, 1, MaxProduct(nums))
	require.Equal(t, 1, maxProduct(nums))

	nums = []int{2, -5, -2, -4, 3}
	require.Equal(t, 24, MaxProduct(nums))
	require.Equal(t, 24, maxProduct(nums))

	nums = []int{-1, -2, -9, -6}
	require.Equal(t, 108, MaxProduct(nums))
	require.Equal(t, 108, maxProduct(nums))

	nums = []int{1, 0, -1, 2, 3, -5, -2}
	require.Equal(t, 60, MaxProduct(nums))
	require.Equal(t, 60, maxProduct(nums))

	nums = []int{-2, 1, 0, -3}
	require.Equal(t, 1, MaxProduct(nums))
	require.Equal(t, 1, maxProduct(nums))

	nums = []int{2, -1, -1, 2, 0, -3, 3}
	require.Equal(t, 4, MaxProduct(nums))
	require.Equal(t, 4, maxProduct(nums))
}
