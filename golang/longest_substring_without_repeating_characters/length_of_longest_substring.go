package longest_substring_without_repeating_characters

func LengthOfLongestSubstring(s string) int {
	lookup := make(map[byte]int)

	var length, startIdx int
	for i := range s {
		if _, ok := lookup[s[i]]; ok && startIdx <= lookup[s[i]] {
			startIdx = lookup[s[i]] + 1
		} else {
			length = Max(length, i - startIdx + 1)
		}

		lookup[s[i]] = i
	}
	return length
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
