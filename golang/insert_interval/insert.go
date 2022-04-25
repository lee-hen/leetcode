package insert_interval

func Insert(intervals [][]int, newInterval []int) [][]int {
	// init data
	newStart, newEnd := newInterval[0], newInterval[1]
	idx, n := 0, len(intervals)
	output := make([][]int, 0)

	// add all intervals starting before newInterval
	for idx < n && newStart > intervals[idx][0] {
		output = append(output, intervals[idx])
		idx++
	}

	// add newInterval
	interval := make([]int, 2)
	// if there is no overlap, just add the interval
	if len(output) == 0 || output[len(output)-1][1] < newStart {
		output = append(output, newInterval)
	} else { // if there is an overlap, merge with the last interval
		interval, output = output[len(output)-1], output[:len(output)-1]
		interval[1] = Max(interval[1], newEnd)
		output = append(output, interval)
	}

	// add next intervals, merge with newInterval if needed
	for idx < n {
		interval, idx = intervals[idx], idx+1
		start, end := interval[0], interval[1]

		// if there is no overlap, just add an interval
		if output[len(output)-1][1] < start {
			output = append(output, interval)
		} else { // if there is an overlap, merge with the last interval
			interval, output = output[len(output)-1], output[:len(output)-1]
			interval[1] = Max(interval[1], end)
			output = append(output, interval)
		}
	}
	return output
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var res = make([][]int, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	idx1 := BinarySearch(intervals, newInterval[0])
	idx2 := BinarySearch(intervals, newInterval[1])

	if newInterval[0] < intervals[idx1][0] {
		res = append(res, intervals[:idx1]...)
		if newInterval[1] < intervals[idx2][0] {
			res = append(res, newInterval)
			res = append(res, intervals[idx2:]...)
		}  else if newInterval[1] >= intervals[idx2][0] && newInterval[1] <= intervals[idx2][1] {
			intervals[idx2][0] = Min(intervals[idx2][0], newInterval[0])
			intervals[idx2][1] = Max(intervals[idx2][1], newInterval[1])
			res = append(res, intervals[idx2:]...)
		} else {
			res = append(res, newInterval)
			if idx2 < len(intervals)-1 {
				res = append(res, intervals[idx2+1:]...)
			}
		}
	} else if newInterval[0] >= intervals[idx1][0] && newInterval[0] <= intervals[idx1][1] {
		intervals[idx1][0] = Min(intervals[idx1][0], newInterval[0])
		intervals[idx1][1] = Max(intervals[idx1][1], newInterval[0])

		if newInterval[1] < intervals[idx2][0] {
			intervals[idx1][1] = Max(intervals[idx1][1], newInterval[1])
			res = append(res, intervals[:idx1+1]...)
			res = append(res, intervals[idx2:]...)
		}  else if newInterval[1] >= intervals[idx2][0] && newInterval[1] <= intervals[idx2][1] {
			intervals[idx1][1] = Max(intervals[idx2][1], newInterval[1])
			res = append(res, intervals[:idx1+1]...)
			if idx2 < len(intervals)-1 {
				res = append(res, intervals[idx2+1:]...)
			}
		} else {
			intervals[idx1][1] = Max(intervals[idx1][1], newInterval[1])
			res = append(res, intervals[:idx1+1]...)
			if idx2 < len(intervals)-1 {
				res = append(res, intervals[idx2+1:]...)
			}
		}
	} else {
		res = append(res, intervals[:idx1+1]...)
		if newInterval[1] >= intervals[idx2][0] && newInterval[1] <= intervals[idx2][1] {
			newInterval[1] = Max(intervals[idx2][1], newInterval[1])
			res = append(res, newInterval)
		} else {
			res = append(res, newInterval)
		}

		if idx2 < len(intervals)-1 {
			res = append(res, intervals[idx2+1:]...)
		}
	}

	return res
}

func BinarySearch(intervals [][]int, x int) int {
	low := 0
	high := len(intervals) - 1
	var mid int
	for low <= high {
		mid = low + (high - low) / 2

		if intervals[mid][1] < x {
			if mid == len(intervals)-1 || x < intervals[mid+1][0] {
				return mid
			}

			low = mid + 1
		} else if intervals[mid][0] > x {
			if mid == 0 || x > intervals[mid-1][1] {
				return mid
			}

			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
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
