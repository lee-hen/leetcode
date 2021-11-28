package product_of_array_except_self

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1,2,3,4}
	require.Equal(t, []int{24,12,8,6}, productExceptSelf(nums))
	require.Equal(t, []int{24,12,8,6}, ProductExceptSelf(nums))
}
