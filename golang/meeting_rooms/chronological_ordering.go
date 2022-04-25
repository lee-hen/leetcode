package meeting_rooms

import "sort"

func ChronologicalOrderingMeeting(intervals [][]int) int  {
	if len(intervals) == 0 {
		return 0
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


	// When we encounter an ending event, that means that some meeting that started earlier has ended now. We are not really concerned with which meeting has ended.
	// All we need is that some meeting ended thus making a room available.
	rooms := 0
	var i, j int
	for i < len(intervals) {
		if startTimings[i] < endTimings[j] {
			rooms++
			i++
		} else {
			i++
			j++
		}
	}

	return rooms
}
