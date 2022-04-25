package meeting_rooms

import "sort"

func CanAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	startTimings := make([]int, 0)
	endTimings := make([]int, 0)

	for _, interval := range intervals {
		startTimings = append(startTimings, interval[0])
		endTimings = append(endTimings, interval[1])
	}

	sort.Slice(startTimings, func(i, j int) bool {
		return startTimings[i] < startTimings[j]
	})

	sort.Slice(endTimings, func(i, j int) bool {
		return endTimings[i] < endTimings[j]
	})

	var i, j = 1, 0
	for i < len(intervals) {
		if startTimings[i] < endTimings[j] {
			return false
		} else {
			i++
			j++
		}
	}

	return true
}
