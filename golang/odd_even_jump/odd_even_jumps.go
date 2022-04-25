package odd_even_jump

import (
	"sort"
)

type item struct {
	val int
	idx int
}

func OddEvenJumps(arr []int) int {
	var items []item
	for i, v := range arr {
		items = append(items, item{v, i})
	}

	sort.Slice(items, func (i, j int) bool {
		if items[i].val != items[j].val {
			return items[i].val < items[j].val
		}
		return items[i].idx < items[j].idx
	})

	nextHigher := make([]int, len(arr))
	buildNextJump(items, nextHigher)

	sort.Slice(items, func (i, j int) bool {
		if items[i].val != items[j].val {
			return items[i].val > items[j].val
		}
		return items[i].idx < items[j].idx
	})

	nextLower := make([]int, len(arr))
	buildNextJump(items, nextLower)

	odd, even := make([]bool, len(arr)), make([]bool, len(arr))
	odd[len(arr)-1], even[len(arr)-1] = true, true

	for i := len(arr)-2; i >= 0; i-- {
		odd[i] = even[nextHigher[i]]
		even[i] = odd[nextLower[i]]
	}

	return sum(odd)
}

func buildNextJump(items []item, nextJump []int) {
	var stack []int
	for _, v := range items {
		for len(stack) > 0 && stack[len(stack)-1] < v.idx {
			nextJump[stack[len(stack)-1]] = v.idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v.idx)
	}
}

func sum(arr []bool) int {
	cnt := 0
	for _, num := range arr {
		if num {
			cnt++
		}
	}

	return cnt
}

func oddEvenJumps(arr []int) int {
	oddIndices := make([]int, 0)
	evenIndices := make([]int, 0)

	for i := range arr {
		oddIdx, enenIdx := findNextJump(i, arr)
		oddIndices = append(oddIndices, oddIdx)
		evenIndices = append(evenIndices, enenIdx)
	}

	doJump := func(arr, oddIndices, evenIndices []int) int{
		var jumps int
		for i := range arr {
			var jump = 1
			var jumpIdx = i
			for {
				if jump & 1 == 0 {
					jumpIdx = evenIndices[jumpIdx]
				} else {
					jumpIdx = oddIndices[jumpIdx]
				}

				if jumpIdx == len(arr)-1 {
					jumps++
					break
				}

				if jumpIdx == -1 {
					break
				}
				jump++
			}
		}

		return jumps
	}

	return doJump(arr, oddIndices, evenIndices)
}

func findNextJump(i int, arr []int) (int, int) {
	if i == len(arr)-1 {
		return i, i
	}

	var higher, lower int
	var enenIdx = -1
	var oddIdx = -1
	for j := i+1; j < len(arr); j++ {
		if enenIdx == -1 {
			if arr[j] <= arr[i] {
				lower = arr[j]
				enenIdx = j
			}
		} else if arr[j] <= arr[i] && arr[j] > lower {
			lower = arr[j]
			enenIdx = j
		}

		if oddIdx == -1 {
			if arr[j] >= arr[i] {
				higher = arr[j]
				oddIdx = j
			}
		} else if arr[j] >= arr[i] && arr[j] < higher {
			higher = arr[j]
			oddIdx = j
		}
	}

	return oddIdx, enenIdx
}
