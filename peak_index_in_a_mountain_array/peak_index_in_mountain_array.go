package peak_index_in_a_mountain_array

func PeakIndexInMountainArray(arr []int) int {
	lo,  hi := 0, len(arr)-1

	for lo < hi {
		mi := lo + (hi - lo) / 2
		if arr[mi] < arr[mi + 1] {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	return lo
}

func peakIndexInMountainArray(arr []int) int {
	curr := arr[0]
	var idx int
	for i, num := range arr {
		if curr < num {
			curr = num
			idx = i
		}
	}
	return idx
}
