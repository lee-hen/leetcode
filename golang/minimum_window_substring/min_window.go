package minimum_window_substring

// https://github.com/lee-hen/ctci/ch_17_hard/q17_18_shortest_supersequence

type CountedLookup struct {
	lookup        map[byte]int
	runningLookup map[byte]int
	fulfilled     int
}

func NewCountedLookup(array []byte) *CountedLookup {
	c := new(CountedLookup)
	c.lookup = make(map[byte]int)
	c.runningLookup = make(map[byte]int)
	for _, a := range array {
		c.lookup[a]++
	}

	return c
}

func (c *CountedLookup) Contains (v byte) bool {
	_, ok := c.lookup[v]
	return ok
}

func (c *CountedLookup) IncrementIfFound(v byte) {
	if !c.Contains(v) {
		return
	}

	c.runningLookup[v]++
	if c.lookup[v] == c.runningLookup[v] {
		c.fulfilled++
	}
}

func (c *CountedLookup) DecrementIfFound(v byte) {
	if !c.Contains(v) {
		return
	}

	if c.runningLookup[v] == c.lookup[v] {
		c.fulfilled--
	}
	c.runningLookup[v]--
}

func (c *CountedLookup) AreAllFulfilled() bool {
	return c.fulfilled == len(c.lookup)
}

type Pair struct {
	idx int
	char byte
}

func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	countedLookup := NewCountedLookup([]byte(t))

	filteredStrs := make([]Pair, 0)
	for i := 0; i < len(s); i++ {
		if countedLookup.Contains(s[i]) {
			filteredStrs = append(filteredStrs, Pair{i, s[i]})
		}
	}

	if len(filteredStrs) == 0 {
		return ""
	}

	var j int
	var best *Range
	countedLookup.IncrementIfFound(filteredStrs[0].char)
	for i := range filteredStrs {
		left := filteredStrs[i].idx
		j = findClosure(&filteredStrs, j, countedLookup)
		right := filteredStrs[j].idx

		if !countedLookup.AreAllFulfilled() {
			break
		}

		length := right - left + 1
		if best == nil || best.Length() > length {
			best = NewRange(left, right)
		}

		countedLookup.DecrementIfFound(filteredStrs[i].char)
	}

	if best == nil {
		return ""
	}

	return s[best.start:best.end+1]
}

func findClosure(filteredStrs *[]Pair, startIndex int, countedLookup *CountedLookup) int {
	pairs := *filteredStrs
	index := startIndex

	for !countedLookup.AreAllFulfilled() && index < len(pairs)-1 {
		index++
		countedLookup.IncrementIfFound(pairs[index].char)
	}
	return index
}

type Range struct {
	start int
	end int
}

func NewRange(s, e int) *Range {
	r := new(Range)
	r.start = s
	r.end = e

	return r
}

func (r *Range) Length() int {
	return r.end - r.start + 1
}
