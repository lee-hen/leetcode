package next_permutation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 5, 2, 4, 3}
	NextPermutation(nums)
	require.Equal(t, []int{1, 5, 3, 2, 4}, nums)

	nums = []int{3, 2, 1}
	NextPermutation(nums)
	require.Equal(t, []int{1, 2, 3}, nums)

	nums = []int{1, 1, 5}
	NextPermutation(nums)
	require.Equal(t, []int{1, 5, 1}, nums)
}
