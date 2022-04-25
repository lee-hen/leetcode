package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i - 1] * answer[i - 1]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * right
		right *= nums[i]
	}

	return answer
}

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	for i := range left {
		left[i] = 1
		right[i] = 1
	}

	for i, j := 1, len(nums)-2; i < len(nums) && j >= 0; i, j = i+1, j-1 {
		left[i] *= left[i-1] * nums[i-1]
		right[j] *= right[j+1] * nums[j+1]
	}

	exceptSelf := make([]int, len(nums))
	for i := range exceptSelf {
		exceptSelf[i] = left[i] * right[i]
	}

	return exceptSelf
}
