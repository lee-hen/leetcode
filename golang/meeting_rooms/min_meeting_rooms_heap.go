package meeting_rooms

import (
	"container/heap"
	"sort"
)

func MinMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	intHeap := make(IntHeap, 1)
	intHeap[0] = intervals[0][1]
	meetingRooms := &intHeap
	heap.Init(meetingRooms)

	var rooms = 1
	for i := 1; i < len(intervals); i++ {
		currentEnd := (*meetingRooms)[0]

		if intervals[i][0] < currentEnd {
			rooms++
			heap.Push(meetingRooms, intervals[i][1])
		} else {
			(*meetingRooms)[0] = intervals[i][1]
			heap.Fix(meetingRooms, 0)
		}
	}

	return rooms
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
