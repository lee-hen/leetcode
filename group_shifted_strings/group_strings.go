package group_shifted_strings

import "fmt"

func GroupStrings(strings []string) [][]string {
	stringsMap := make(map[string][]string)
	for _, s := range strings {
		key := ""
		for i := 1; i < len(s); i++ {
			asciiCodeCircularDiff := 26 + s[i] - s[i-1]
			key += generateKey(asciiCodeCircularDiff % 26)
		}
		stringsMap[key] = append(stringsMap[key], s)
	}

	groups := make([][]string, 0)
	for _, letters := range stringsMap {
		groups = append(groups, letters)
	}
	return groups
}

func generateKey(diff uint8) string {
	return fmt.Sprintf("-%d-", diff)
}

func GroupStringsBruteForce(strings []string) [][]string {
	stringsMap := make(map[int][]string)

	for _, str := range strings {
		stringsMap[len(str)] = append(stringsMap[len(str)], str)
	}

	groups := make([][]string, 0)
	for key, letters := range stringsMap {
		length := key

		var startGroupIdx = -1
		group, isInGroups := &[]string{letters[0]}, false

		for i := 1; i < len(letters); i++ {
			if !couldFormSequence(letters[i-1], letters[i], length) {
				if !isInGroups {
					groups = append(groups, *group)

					if startGroupIdx == -1 {
						startGroupIdx = len(groups)-1
					}
				}

				isInGroups = false
				for groupIdx := startGroupIdx; groupIdx <= len(groups)-1; groupIdx++ {
					if couldFormSequence(groups[groupIdx][0], letters[i], length) {
						group = &groups[groupIdx]
						isInGroups = true
						break
					}
				}

				if !isInGroups {
					group = &[]string{}
				}
			}

			*group = append(*group, letters[i])
		}

		if !isInGroups {
			groups = append(groups, *group)
		}
	}
	return groups
}

func couldFormSequence(str1, str2 string, length int) bool {
	if str2[0] < str1[0] {
		return couldFormSequence(str2, str1, length)
	}
	asciiCodeDistance := str2[0] - str1[0]
	for j := 1; j < length; j++ {
		curr := str2[j] - 97
		prev := str1[j] - 97

		if (prev + asciiCodeDistance) % 26 != curr {
			return false
		}
	}

	return true
}
