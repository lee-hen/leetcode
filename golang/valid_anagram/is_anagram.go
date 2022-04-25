package valid_anagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[byte]int)
	for i := range s {
		c := s[i]
		freq[c]++
	}

	for j := range t {
		c := t[j]
		freq[c]--
		if freq[c] < 0 {
			return false
		}
	}

	return true
}
