package three_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	require.Equal(t, [][]int{{-1,-1,2}, {-1,0,1}}, ThreeSum(nums))
}
