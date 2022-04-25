package median_of_two_sorted_arrays

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	var lo int
	hi := len(nums1)

	for lo <= hi {
		mid1 := (lo + hi) / 2
		mid2 := (len(nums1) + len(nums2) + 1) / 2 - mid1

		var maxLeft1, maxLeft2, minRight1, minRight2 int
		if mid1 == 0 {
			maxLeft1 = math.MinInt32
		} else {
			maxLeft1 = nums1[mid1-1]
		}

		if mid1 ==  len(nums1) {
			minRight1 = math.MaxInt32
		} else {
			minRight1 = nums1[mid1]
		}

		if mid2 == 0 {
			maxLeft2 = math.MinInt32
		} else {
			maxLeft2 = nums2[mid2-1]
		}

		if mid2 == len(nums2) {
			minRight2 = math.MaxInt32
		} else {
			minRight2 = nums2[mid2]
		}

		if maxLeft1 > minRight2 {
			hi = mid1-1
		} else if maxLeft2 > minRight1 {
			lo = mid1+1
		} else {
			if (len(nums1) + len(nums2)) & 1 == 0 {
				return (math.Max(float64(maxLeft1), float64(maxLeft2)) + math.Min(float64(minRight1), float64(minRight2))) / 2
			} else {
				return math.Max(float64(maxLeft1), float64(maxLeft2))
			}
		}
	}
	return -1
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	mid := n / 2
	var idx1, idx2 int
	var half = make([]int, 0)
	for idx1 < len(nums1) && idx2 < len(nums2) && len(half)-1 < mid {
		if nums1[idx1] <= nums2[idx2] {
			half = append(half, nums1[idx1])
			idx1++
		} else {
			half = append(half, nums2[idx2])
			idx2++
		}
	}

	for idx1 < len(nums1) && len(half)-1 < mid {
		half = append(half, nums1[idx1])
		idx1++
	}

	for idx2 < len(nums2) && len(half)-1 < mid {
		half = append(half, nums2[idx2])
		idx2++
	}

	var ans float64
	if n & 1 == 0 {
		ans = float64(half[len(half)-1] + half[len(half)-2]) / 2.0
	} else {
		ans = float64(half[len(half)-1])
	}

	return ans
}
