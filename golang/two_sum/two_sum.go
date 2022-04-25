package two_sum

func twoSum(nums []int, target int) []int {
	memorize := make(map[int]int)

	for i, y := range nums {
		x := target - y

		if j, ok := memorize[x]; ok {
			return []int{i, j}
		}

		memorize[y] = i
	}
	return []int{}
}
