package find_and_replace_in_string

import (
	"sort"
	"strings"
)

func FindReplaceString(s string, indices []int, sources []string, targets []string) string {
	inverseIndices := make(map[int]int)
	for i := range indices {
		inverseIndices[indices[i]] = i
	}

	sort.Ints(indices)
	builder := new(strings.Builder)

	var i, idx, prevIdx int
	for i < len(indices) {
		index := indices[i]
		if prevIdx > index {
			return ""
		}

		if idx == len(s) {
			break
		}

		if j := inverseIndices[index]; index == idx  {
			if sources[j] == s[index: index+len(sources[j])] {
				builder.WriteString(targets[j])
				idx = index+len(sources[j])
				prevIdx = idx
			}
			i++
		} else {
			builder.WriteByte(s[idx])
			idx++
		}
	}

	builder.WriteString(s[idx:])
	return builder.String()
}
