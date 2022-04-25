package meeting_rooms

import (
	"math"
	"sort"
)

// [2 15] [4 9] [9 29] [16 23] [36 45]
// 1: 2           15 16     23 36      45
// 2:   4    9 9        29

func MinMeetingRoomsBruteForce(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var rooms = 1
	roomIdx := make(map[int]int)
	meetingRooms := make(map[int]int)
	meetingRooms[rooms] = intervals[0][1]
	roomIdx[intervals[0][1]] = rooms

	for i := 1; i < len(intervals); i++ {
		currentEnd := minMeetingEnd(meetingRooms)

		if intervals[i][0] < currentEnd {
			rooms++
			meetingRooms[rooms] = intervals[i][1]
			roomIdx[intervals[i][1]] = rooms
		} else if idx, ok := roomIdx[currentEnd]; ok  {
			meetingRooms[idx] = intervals[i][1]
			roomIdx[intervals[i][1]] = idx
		}
	}

	return rooms
}

func minMeetingEnd(meetingRooms map[int]int) int {
	m := math.MaxInt32
	for _, end := range meetingRooms {
		m = min(m, end)
	}
	return m
}

func min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
