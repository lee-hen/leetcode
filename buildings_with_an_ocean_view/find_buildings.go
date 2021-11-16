package buildings_with_an_ocean_view

func FindBuildings(heights []int) []int {
	oceanView := make([]int, 0)
	oceanView = append(oceanView, len(heights)-1)
	max := heights[len(heights)-1]
	for i := len(heights)-2; i >= 0; i-- {
		if heights[i] > max {
			max = heights[i]
			oceanView = append(oceanView, i)
		}
	}

	for i, j := 0, len(oceanView)-1; i < j; i, j = i+1, j-1 {
		oceanView[i], oceanView[j] = oceanView[j], oceanView[i]
	}

	return oceanView
}
