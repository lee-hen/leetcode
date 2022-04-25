package three_sum

import "sort"

type triplet struct {
	num1, num2, num3 int
}

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := make([][]int, 0)
	seen := make(map[triplet]bool)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			t := triplet{nums[i], nums[left], nums[right]}

			if currentSum == 0 {
				if !seen[t] {
					triplets = append(triplets, []int{nums[i], nums[left], nums[right]})
					seen[t] = true
				}
				left++
				right -= 1
			} else if currentSum < 0 {
				left++
			} else if currentSum > 0 {
				right--
			}
		}
	}
	return triplets
}

