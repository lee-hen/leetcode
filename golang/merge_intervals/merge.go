package merge_intervals

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	for _, curr := range intervals {
		if len(res) > 0 {
			prev := res[len(res)-1]
			if prev[1] < curr[0] {
				res = append(res, curr)
			} else {
				if prev[1] >= curr[0] && prev[1] <= curr[1] {
					res[len(res)-1][1] = Max(prev[1], curr[1])
				}
			}
		} else {
			res = append(res, curr)
		}
	}
	return res
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
