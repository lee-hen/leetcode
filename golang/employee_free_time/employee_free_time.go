package employee_free_time

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func EmployeeFreeTime(schedule [][]*Interval) []*Interval {
	ints := make([]*Interval, 0)
	for _, intervals := range schedule {
		for _, interval := range intervals {
			ints = append(ints, interval)
		}
	}

	sort.Slice(ints, func(i, j int) bool {
		return ints[i].Start < ints[j].Start
	})

	res := make([]*Interval, 0)
	end := ints[0].End

	for i := 1; i < len(ints); i++ {
		if ints[i].Start > end {
			res = append(res, &Interval{end, ints[i].Start})
		}

		end = max(end, ints[i].End)
	}

	return res
}

func max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	startTimings := make([]int, 0)
	endTimings := make([]int, 0)

	for _, intervals := range schedule {
		for i := 0; i < len(intervals); i++ {
			startTimings = append(startTimings, intervals[i].End)
			endTimings = append(endTimings, intervals[i].Start)
		}
	}

	sort.Slice(startTimings, func(i, j int) bool {
		return startTimings[i] < startTimings[j]
	})

	sort.Slice(endTimings, func(i, j int) bool {
		return endTimings[i] < endTimings[j]
	})

	freeTimes := make(map[int]int)
	var i, j int
	for i < len(startTimings) && j < len(endTimings) {
		if startTimings[i] < endTimings[j] {
			freeTimes[endTimings[j]] = startTimings[i]
			i++
		} else {
			j++
		}
	}

	commonFreeIntervals := make([]*Interval, 0)
	for end, start := range freeTimes {
		if onWorkTime(end, schedule) {
			continue
		}

		commonFreeIntervals = append(commonFreeIntervals, &Interval{start, end})
	}

	sort.Slice(commonFreeIntervals, func(i, j int) bool {
		return commonFreeIntervals[i].Start < commonFreeIntervals[j].Start
	})

	return commonFreeIntervals
}

func onWorkTime(endTime int, schedule [][]*Interval) bool {
	for _, workTime := range schedule {
		for k := 0; k < len(workTime) && endTime > workTime[k].Start; k++ {
			if workTime[k].Start < endTime && workTime[k].End >= endTime {
				return true
			}
		}
	}

	return false
}
