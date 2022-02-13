package minimum_area_rectangle

import (
	"fmt"
	"math"
)

func MinAreaRect(points [][]int) int {
	lookup := buildPointLookup(points)
	var min = math.MaxInt32

	for idx, point := range points {
		p2x, p2y := point[0], point[1]

		for prev := 0; prev < idx; prev++ {
			p1x, p1y := points[prev][0], points[prev][1]
			pointsShareValue := p1x == p2x || p1y == p2y

			if pointsShareValue {
				continue
			}

			diagonalPoint1 := fmt.Sprintf("%d-%d", p1x, p2y)
			diagonalPoint2 := fmt.Sprintf("%d-%d", p2x, p1y)
			diagonalPoint1Exists := lookup[diagonalPoint1]
			diagonalPoint2Exists := lookup[diagonalPoint2]

			if diagonalPoint1Exists && diagonalPoint2Exists {
				min = Min(min, Abs(p2x-p1x) * Abs(p2y-p1y))
			}
		}
	}
	if min != math.MaxInt32 {
		return min
	}
	return 0
}

func buildPointLookup(points [][]int) map[string]bool {
	lookup := map[string]bool{}
	for _, point := range points {
		x, y := point[0], point[1]
		pointString := fmt.Sprintf("%d-%d", x, y)
		lookup[pointString] = true
	}
	return lookup
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

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
