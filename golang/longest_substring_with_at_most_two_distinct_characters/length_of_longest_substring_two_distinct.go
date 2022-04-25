package longest_substring_with_at_most_two_distinct_characters

func LengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) == 1 {
		return 1
	}

	var a, b byte
	var length, running, continuousLength int
	for i := range s {
		if a == 0 {
			a = s[i]
		}

		if b == 0 && s[i] != a {
			b = s[i]
		}
		if s[i] == a || s[i] == b {
			running++
			if a == 0 || b == 0 {
				continue
			}

			if s[i] == s[i-1] {
				continuousLength++
			} else {
				continuousLength = 1
			}
		} else {
			length = Max(length, running)
			running = continuousLength + 1
			continuousLength = 1
			a = s[i-1]
			b = s[i]
		}
	}

	length = Max(length, running)
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
