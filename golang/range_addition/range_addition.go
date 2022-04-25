package range_addition

func getModifiedArray(length int, updates [][]int) []int {
	array := make([]int, length, length)

	for _, update := range updates {
		for startIdx := update[0]; startIdx <= update[1]; startIdx++ {
			array[startIdx] += update[2]
		}
	}

	return array
}








