package median_of_two_sorted_arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{1}
	require.Equal(t, findMedianSortedArrays(nums1, nums2), FindMedianSortedArrays(nums1, nums2))

	nums1 = []int{2, 2, 4, 4}
	nums2 = []int{2, 2, 4, 4}
	require.Equal(t, findMedianSortedArrays(nums1, nums2), FindMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3}
	nums2 = []int{2, 7}
	require.Equal(t, findMedianSortedArrays(nums1, nums2), FindMedianSortedArrays(nums1, nums2))

	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	require.Equal(t, findMedianSortedArrays(nums1, nums2), FindMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3}
	nums2 = []int{2}
	require.Equal(t, findMedianSortedArrays(nums1, nums2), FindMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	require.Equal(t, findMedianSortedArrays(nums1, nums2), FindMedianSortedArrays(nums1, nums2))
}
