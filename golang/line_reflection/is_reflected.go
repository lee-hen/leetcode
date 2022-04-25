package line_reflection

import (
	"fmt"
	"math"
	"sort"
)

func IsReflected(points [][]int) bool {
	max := math.MinInt32
	min := math.MaxInt32

	set := make(map[string]struct{})

	for _, p := range points {
		max = Max(max, p[0])
		min = Min(min, p[0])
	}

	for _, p := range points{
		key := fmt.Sprintf("%d-%d", p[0], p[1])
		set[key] = struct{}{}
	}

	sum := max+min // p1(x)-min = max - p2(x) => p1(x) + p2(x) = min + max
	for _, p:= range points {
		str := fmt.Sprintf("%d-%d", sum-p[0], p[1]) // sum - p1 = p2
		if _, ok := set[str]; !ok {
			return false
		}
	}
	return true
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

func IsReflected1(points [][]int) bool {
	points = getUniqSortedPoints(points)

	parallelXAxis := make(map[int][]int)
	for i := range points {
		parallelXAxis[points[i][1]] = append(parallelXAxis[points[i][1]], points[i][0])
	}

	var middleXAxis *float64
	for y := range parallelXAxis {
		middle, ok := isSymmetric(parallelXAxis[y])
		if !ok {
			return false
		}

		if middleXAxis == nil {
			middleXAxis = &middle
		} else if *middleXAxis != middle  {
			return false
		}
	}
	return true
}

func isSymmetric(xAxis []int) (float64, bool) {
	var middle *float64
	for i, j := 0, len(xAxis)-1; i <= j; i, j = i+1, j-1 {

		curr := float64(xAxis[i] + xAxis[j]) / 2
		if middle == nil {
			middle = &curr
		} else if *middle != curr  {
			return 0, false
		}
	}
	return *middle, true
}

func getUniqSortedPoints(points [][]int) [][]int {
	var uniqPoints = make([][]int, 0)
	seen := make(map[string]bool)

	for _, point := range points {
		key := fmt.Sprintf("%d, %d", point[0], point[1])
		if !seen[key] {
			uniqPoints = append(uniqPoints, point)
			seen[key] = true
		}
	}

	sort.Slice(uniqPoints, func(i, j int) bool {
		return uniqPoints[i][0] < uniqPoints[j][0]
	})
	return uniqPoints
}
