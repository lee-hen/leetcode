package design_hit_counter

const fiveMinutes = 300

type HitCounter struct {
	hits map[int]int

	lastHitAt int
	fiveMinuteSplitNumber int
}

func Constructor() HitCounter {
	h := new(HitCounter)
	h.hits = make(map[int]int)
	h.lastHitAt = 0

	h.fiveMinuteSplitNumber = 0
	return *h
}

func (h *HitCounter) Hit(timestamp int)  {
	if timestamp / fiveMinutes > h.fiveMinuteSplitNumber + 1 {
		h.hits = make(map[int]int)
	}

	h.hits[timestamp] = h.hits[h.lastHitAt] + 1
	h.lastHitAt = timestamp
	h.fiveMinuteSplitNumber = timestamp / fiveMinutes
}

func (h *HitCounter) GetHits(timestamp int) int {
	if timestamp < h.lastHitAt {
		panic("not allowed")
	}

	if timestamp / fiveMinutes > h.fiveMinuteSplitNumber + 1 {
		return 0
	}

	var target int
	for time := range h.hits {
		if time <= timestamp - fiveMinutes {
			target = Max(time, target)
		}
	}

	return h.hits[h.lastHitAt] - h.hits[target]
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
