package analyze_user_website_visit_pattern

import (
	"sort"
	"strings"
)

func MostVisitedPattern(usernames []string, timestamps []int, websites []string) []string {
	usernames, websites = getSorted(usernames, timestamps,  websites)
	timestampIndices := getTimeStampIndices(timestamps)

	timestampsByName := make(map[string][]int)
	for i := range usernames {
		timestampsByName[usernames[i]] = append(timestampsByName[usernames[i]], timestamps[i])
	}

	timestampPatterns := getTimestampPatterns(timestampsByName)

	patterns := make(map[string]int)
	users := make(map[string]string)

	for _, timestampsPattern := range timestampPatterns {
		timestamp1, timestamp2, timestamp3 := timestampsPattern[0], timestampsPattern[1], timestampsPattern[2]
		idx1, idx2, idx3 := timestampIndices[timestamp1], timestampIndices[timestamp2], timestampIndices[timestamp3]
		patternString := websites[idx1] + "," + websites[idx2] + "," + websites[idx3]

		if users[patternString] != usernames[idx1] {
			patterns[patternString]++
			users[patternString] = usernames[idx1]
		}
	}

	mostVisited := -1
	var patternString string
	var pattern []string
	for patternStr, score := range patterns {
		if score > mostVisited || score == mostVisited && patternString != "" && patternStr < patternString  {
			mostVisited = score
			patternString = patternStr
			pattern = strings.Split(patternString, ",")
		}
	}

	return pattern
}

func getTimestampPatterns(timestampsByName map[string][]int) [][]int{
	patterns := make([][]int, 0)
	for _, timestamps := range timestampsByName {
		if len(timestamps) < 3 {
			continue
		}

		if len(timestamps) == 3 {
			patterns = append(patterns, timestamps)
		} else {
			patterns = append(patterns, generateTimestampPatterns(timestamps)...)
		}
	}

	return patterns
}

func generateTimestampPatterns(timestamps []int) [][]int {
	patterns := make([][]int, 0)
	for i := 0; i < len(timestamps)-2; i++ {
		for j := i+1; j < len(timestamps)-1; j++ {
			for k := j+1; k < len(timestamps); k++ {
				pattern := []int{timestamps[i], timestamps[j], timestamps[k]}
				patterns = append(patterns, pattern)
			}
		}
	}

	return patterns
}

func getSorted(usernames []string, timestamps []int, websites []string) ([]string, []string){
	timestampIndices := getTimeStampIndices(timestamps)

	sort.Ints(timestamps)
	sortedUsername := make([]string, len(timestamps))
	sortedWebsites := make([]string, len(timestamps))
	for i, timestamp := range timestamps {
		sortedUsername[i] = usernames[timestampIndices[timestamp]]
		sortedWebsites[i] = websites[timestampIndices[timestamp]]
	}

	return sortedUsername, sortedWebsites
}

func getTimeStampIndices(timestamps []int) map[int]int {
	timestampIndices := make(map[int]int)
	for i, timestamp := range timestamps {
		timestampIndices[timestamp] = i
	}

	return timestampIndices
}
