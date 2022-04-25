package k_closest_points_to_origin

import (
	"math"
	"sort"
)

func KClosest(points [][]int, k int) [][]int {
	// Precompute the Euclidean distance for each point,
	// define the initial binary search range,
	// and create a reference list of point indices
	distances := make([]float64, len(points))
	low, high := 0.0, 0.0
	remaining := make([]int, 0)

	for i := range points {
		distances[i] = float64(euclideanDistance(points[i]))
		high = math.Max(high, distances[i])
		remaining = append(remaining, i)
	}

	// Perform a binary search of the distances
	// to find the k closest points
	closest := make([]int, 0)
	for k > 0 {
		mid := low + (high - low) / 2
		result := splitDistances(remaining, distances, mid)
		closer := *result[0]
		farther := *result[1]
		if len(closer) > k {
			// If more than k points are in the closer distances
			// then discard the farther points and continue
			remaining = closer
			high = mid
		} else {
			// Add the closer points to the answer array and keep
			// searching the farther distances for the remaining points
			k -= len(closer)
			closest = append(closest, closer...)
			remaining = farther
			low = mid
		}
	}

	// Return the k closest points using the reference indices
	k = len(closest)
	answer := make([][]int, k)
	for i := 0; i < k; i++ {
		answer[i] = points[closest[i]]
	}
	return answer
}

func splitDistances(remaining []int, distances []float64, mid float64) []*[]int {
	// Split the distances around the midpoint
	// and return them in separate lists
	result := make([]*[]int, 0)
	closer := make([]int, 0)
	farther := make([]int, 0)
	result = append(result, &closer)
	result = append(result, &farther)
	for _, point := range remaining {
		if distances[point] <= mid {
			closer = append(closer, point)
		} else {
			farther = append(farther, point)
		}
	}
	return result
}

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		d1 := euclideanDistance(points[i])
		d2 := euclideanDistance(points[j])

		return d1 < d2
	})

	res := make([][]int, 0)
	for i := 1; i <= k; i++ {
		res = append(res, points[i-1])
	}

	return res
}

func euclideanDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
