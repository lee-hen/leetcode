package split_array_largest_sum

import (
	"fmt"
	"math"
)

func SplitArray(nums []int, m int) int {
	// Find the sum of all elements and the maximum element
	sum := 0
	maxElement := math.MinInt32
	maxElement = Max(maxElement, nums...)

	// Define the left and right boundary of binary search
	left := maxElement
	right := sum
	var minimumLargestSplitSum int
	for left <= right {
		// Find the mid value
		maxSumAllowed := (left + right) / 2

		// Find the minimum splits. If splitsRequired is less than
		// or equal to m move towards left i.e., smaller values
		if minimumSubArraysRequired(nums, maxSumAllowed) <= m {
			right = maxSumAllowed - 1
			minimumLargestSplitSum = maxSumAllowed
		} else {
			// Move towards right if splitsRequired is more than m
			left = maxSumAllowed + 1
		}
	}

	return minimumLargestSplitSum
}

func minimumSubArraysRequired(nums []int, maxSumAllowed int) int{
	currentSum := 0
	splitsRequired := 0

	for _, element := range nums {
		// Add element only if the sum doesn't exceed maxSumAllowed
		if currentSum + element <= maxSumAllowed {
			currentSum += element
		} else {
			// If the element addition makes sum more than maxSumAllowed
			// Increment the splits required and reset sum
			currentSum = element
			splitsRequired++
		}
	}

	// Return the number of subArrays, which is the number of splits + 1
	return splitsRequired + 1
}

func splitArray(nums []int, m int) int {
	cache := make(map[string]int)
	return split(0, m, nums, Sum(nums), cache)
}

func split(idx, k int, nums []int, lastSum int, cache map[string]int) int {
	key := fmt.Sprintf("%d-%d", k, idx)

	if _, ok := cache[key]; ok {
		return cache[key]
	}

	if k == 1 {
		cache[key] = lastSum
		return cache[key]
	}

	var eachSum int
	var minimizeLarge = math.MaxInt32
	for i := idx; i <= len(nums)-k; i++ {
		eachSum += nums[i]
		minimizeLarge = Min(minimizeLarge, Max(eachSum, split(i+1, k-1, nums, lastSum-eachSum, cache)))

		// my solution without this, hit time Limit
		if eachSum >= minimizeLarge {
			break
		}
	}

	cache[key] = minimizeLarge
	return minimizeLarge
}

func SplitArrayBottomUp(nums []int,  m int) int {
	n := len(nums)
	// Store the prefix sum of nums array
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	memo := make([][]int, 1001)
	for i := range memo {
		memo[i] = make([]int, 51)
	}

	for subarrayCount := 1; subarrayCount <= m; subarrayCount++ {
		for currIndex := 0; currIndex < n; currIndex++ {
			// Base Case: If there is only one subarray left, then all of the remaining numbers
			// must go in the current subarray. So return the sum of the remaining numbers.
			if subarrayCount == 1 {
				memo[currIndex][subarrayCount] = prefixSum[n] - prefixSum[currIndex]
				continue
			}

			// Otherwise, use the recurrence relation to determine the minimum largest subarray
			// sum between currIndex and the end of the array with subarrayCount subarray remaining.
			minimumLargestSplitSum := math.MaxInt32
			for i := currIndex; i <= n - subarrayCount; i++ {
				// Store the sum of the first subarray.
				firstSplitSum := prefixSum[i + 1] - prefixSum[currIndex]

				// Find the maximum subarray sum for the current first split.
				largestSplitSum := Max(firstSplitSum, memo[i + 1][subarrayCount - 1])

				// Find the minimum among all possible combinations.
				minimumLargestSplitSum = Min(minimumLargestSplitSum, largestSplitSum)

				if firstSplitSum >= minimumLargestSplitSum {
					break
				}
			}

			memo[currIndex][subarrayCount] = minimumLargestSplitSum
		}
	}

	return memo[0][m]
}

func SplitArrayTopDown(nums []int,  m int) int {
	// Store the prefix sum of nums array.
	prefixSum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	memo := make(map[string]int)
	return getMinimumLargestSplitSum(prefixSum, 0, m, memo)
}

func getMinimumLargestSplitSum(prefixSum []int, idx int, k int, memo map[string]int) int {
	n := len(prefixSum) - 1
	key := fmt.Sprintf("%d-%d", k, idx)

	// We have already calculated the answer so no need to go into recursion
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	// Base Case: If there is only one subarray left, then all of the remaining numbers
	// must go in the current subarray. So return the sum of the remaining numbers.
	if k == 1 {
		memo[key] = prefixSum[n] - prefixSum[idx]
		return memo[key]
	}

	// Otherwise, use the recurrence relation to determine the minimum largest subarray
	// sum between currIndex and the end of the array with subarrayCount subarrays remaining.
	minimumLargestSplitSum := math.MaxInt32

	for  i := idx; i <= n-k; i++ {
		// Store the sum of the first subarray.
		firstSplitSum := prefixSum[i+1] - prefixSum[idx]

		// Find the maximum subarray sum for the current first split.
		largestSplitSum := Max(firstSplitSum, getMinimumLargestSplitSum(prefixSum, i+1, k-1, memo))

		// Find the minimum among all possible combinations.
		minimumLargestSplitSum = Min(minimumLargestSplitSum, largestSplitSum)

		if firstSplitSum >= minimumLargestSplitSum {
			break
		}
	}

	memo[key] = minimumLargestSplitSum
	return memo[key]
}

func Sum(arr []int) int {
	var res int
	for i := range arr {
		res += arr[i]
	}
	return res
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

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}