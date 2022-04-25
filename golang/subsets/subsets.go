package subsets

func Subsets(num []int) [][]int {
	allSubsets := make([][]int, 0)

	max := 1 << len(num)
	for k := 0; k < max; k++ {
		subset := convertIntToSet(k, num)
		allSubsets = append(allSubsets, subset)
	}
	return allSubsets
}

func convertIntToSet(x int, num []int) []int {
	subset := make([]int, 0)
	index := 0
	for k := x; k > 0; k >>= 1 {
		if (k & 1) == 1 {
			subset = append(subset, num[index])
		}
		index++
	}
	return subset
}
