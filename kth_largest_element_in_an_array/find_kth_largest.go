package kth_largest_element_in_an_array

import (
	"github.com/lee-hen/leecode/lib"
)

type PartitionResult struct {
	leftSize, middleSize int
}

func NewPartitionResult(left, middle int) *PartitionResult{
	return &PartitionResult{
		leftSize: left,
		middleSize: middle,
	}
}

func FindKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return Min(nums[0], nums...)
	}

	return rank(nums, len(nums)-k, 0, len(nums)-1) // the heap solution is more simple
}

func rank(array []int, k, start, end int) int {
	if k >= len(array) {
		return len(array)-1
	}

	/* Partition array around an arbitrary pivot. */
	pivot := array[randomIntInRange(start, end)]
	partition := partition(array, start, end, pivot)
	leftSize := partition.leftSize
	middleSize := partition.middleSize

	if k < leftSize { // Rank k is on left half
		return rank(array, k, start, start + leftSize - 1)
	} else if k < leftSize + middleSize { // Rank k is in middle
		return pivot // middle is all pivot values
	} else { // Rank k is on right
		return rank(array, k - leftSize - middleSize, start + leftSize + middleSize, end)
	}
}

func randomIntInRange(min, max int) int {
	return lib.RandomIntInRange(min, max)
}

func partition(array []int, start, end, pivot int) *PartitionResult{
	left := start /* Stays at (right) edge of left side. */
	right := end  /* Stays at (left) edge of right side. */
	middle := start /* Stays at (right) edge of middle. */
	for middle <= right {
		if array[middle] < pivot {
			/* Middle is smaller than the pivot. Left is either
			 * smaller or equal to the pivot. Either way, swap
			 * them. Then middle and left should move by one.
			 */
			swap(array, middle, left)
			middle++
			left++
		} else if array[middle] > pivot {
			/* Middle is bigger than the pivot. Right could have
			 * any value. Swap them, then we know that the new
			 * right is bigger than the pivot. Move right by one.
			 */
			swap(array, middle, right)
			right--
		} else if array[middle] == pivot {
			/* Middle is equal to the pivot. Move by one. */
			middle++
		}
	}
	/* Return sizes of left and middle. */
	return NewPartitionResult(left - start, right - left + 1)
}

func swap(array []int, i, j int) {
	t := array[i]
	array[i] = array[j]
	array[j] = t
}

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
