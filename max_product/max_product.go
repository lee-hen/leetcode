package max_product

import "math"

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar := nums[0]
	minSoFar := nums[0]
	result := maxSoFar

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		tempMax := Max(curr, Max(maxSoFar * curr, minSoFar * curr))
		minSoFar = Min(curr, Min(maxSoFar * curr, minSoFar * curr))

		maxSoFar = tempMax
		result = Max(maxSoFar, result)
	}

	return result
}

func maxProduct(nums []int) int {
	products := make([]int, len(nums))

	runningProduct := 1
	result := math.MinInt32
	for i := range nums {
		runningProduct *= nums[i]
		products[i] = runningProduct

		result = Max(result, runningProduct)
		if runningProduct == 0 {
			runningProduct = 1
		}
	}

	leftProduct := 1
	for i := range products {
		result = Max(result, products[i] / leftProduct)

		if leftProduct == 1 && products[i] < 0 {
			leftProduct = products[i]
		}

		if products[i] == 0 {
			leftProduct = 1
		}
	}

	return result
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