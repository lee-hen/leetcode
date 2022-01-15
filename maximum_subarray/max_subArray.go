package maximum_subarray

import (
	"math"
)

func MaxSubArray(nums []int) int {
	// Initialize our variables using the first element.
	currentSubarray := nums[0]
	maxSubarray := nums[0]

	// Start with the 2nd element since we already used the first one.
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// If current_subarray is negative, throw it away. Otherwise, keep adding to it.
		currentSubarray = Max(num, currentSubarray + num)
		maxSubarray = Max(maxSubarray, currentSubarray)
	}

	return maxSubarray
}

func maxSubArray(nums []int) int {
	runningSum := 0
	maxSum := math.MinInt32
	for i := range nums {
		runningSum += nums[i]
		maxSum = Max(maxSum, runningSum)
		if runningSum < 0 {
			runningSum = 0
		}
	}

	return maxSum
}

func Max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}