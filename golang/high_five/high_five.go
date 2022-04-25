package high_five

import (
	"sort"
)

func HighFive(items [][]int) [][]int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][1] > items[j][1]
	})

	topFives := make(map[int][]int)

	for _, item := range items {
		if len(topFives[item[0]]) < 5 {
			topFives[item[0]] = append(topFives[item[0]], item[1])
		}
	}

	result := make([][]int, 0)
	for id, topFive := range topFives {
		sum := 0
		for _, score := range topFive {
			sum += score
		}
		result = append(result, []int{id, sum/5})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}
