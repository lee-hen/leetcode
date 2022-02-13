package isomorphic_strings

import (
	"strconv"
	"strings"
)

func IsIsomorphic(s string, t string) bool {
	return transformString(s) == transformString(t)
}

func transformString(s string) string {
	firstOccurrence := make(map[byte]int)
	builder := new(strings.Builder)

	for i := range s {
		c := s[i]

		if _, ok := firstOccurrence[c]; !ok {
			firstOccurrence[c] = i
		}
		builder.WriteString(strconv.Itoa(firstOccurrence[c]))

		builder.WriteString(" ")
	}
	return builder.String()
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	compressStr1 := compress(s)
	compressStr2 := compress(t)

	freq1 := compressStr1.freq
	freq2 := compressStr2.freq
	if len(freq1) != len(freq2) {
		return false
	}

	s1 := compressStr1.compress
	s2 := compressStr2.compress
	if len(s1) != len(s2) {
		return false
	}

	occurrences1 := compressStr1.occurrences
	occurrences2 := compressStr2.occurrences

	for i := range s1 {
		c1, c2 := s1[i], s2[i]
		if freq1[c1] != freq2[c2] {
			return false
		}

		if occurrences1[i] == occurrences2[i] {
			freq1[c1]--
			freq2[c2]--
		}
	}

	return true
}

type CompressString struct {
	occurrences []int

	compress string
	freq map[byte]int
}

func compress(str string) CompressString {
	compressed := new(strings.Builder)
	var countConsecutive int
	occurrences := make([]int, 0)
	freq := make(map[byte]int)

	for i := range str {
		countConsecutive++

		if i+1 >= len(str) || str[i] != str[i+1] {
			occurrences = append(occurrences, countConsecutive)
			compressed.WriteByte(str[i])
			countConsecutive = 0
		}

		freq[str[i]]++
	}

	return CompressString {
		occurrences: occurrences,
		compress: compressed.String(),
		freq: freq,
	}
}
