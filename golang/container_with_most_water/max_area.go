package container_with_most_water

func MaxArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	leftIdx, rightIdx := 0, len(heights) - 1
	surfaceArea := 0
	for leftIdx < rightIdx {
		if heights[leftIdx] < heights[rightIdx] {
			surfaceArea = Max((rightIdx-leftIdx) * heights[leftIdx], surfaceArea)
			leftIdx++
		} else {
			surfaceArea = Max((rightIdx-leftIdx) * heights[rightIdx], surfaceArea)
			rightIdx--
		}
	}

	return surfaceArea
}

func maxArea(heights []int) int {
	surfaceArea := 0
	for i := range heights {
		for j := i+1; j < len(heights); j++ {
			surfaceArea = Max(Min(heights[i], heights[j]) * (j - i), surfaceArea)
		}
	}

	return surfaceArea
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

func Max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}

