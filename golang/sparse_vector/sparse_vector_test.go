package sparse_vector

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	nums1 := []int{1,0,0,2,3}
	nums2 := []int{0,3,0,4,0}
	v1 := Constructor(nums1)
	v2 := Constructor(nums2)

	require.Equal(t, 8, v1.dotProduct(v2))
}

func TestCase2(t *testing.T) {
	nums1 := []int{0,1,0,0,0}
	nums2 := []int{0,0,0,0,2}
	v1 := Constructor(nums1)
	v2 := Constructor(nums2)

	require.Equal(t, 0, v1.dotProduct(v2))
}


func TestCase3(t *testing.T) {
	nums1 := []int{0,1,0,0,2,0,0}
	nums2 := []int{1,0,0,0,3,0,4}
	v1 := Constructor(nums1)
	v2 := Constructor(nums2)

	require.Equal(t, 6, v1.dotProduct(v2))
}
