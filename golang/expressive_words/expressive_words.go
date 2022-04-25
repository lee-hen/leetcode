package expressive_words

import (
	"strings"
)

func ExpressiveWords(s string, words []string) int {
	stretchyStrCnt := make([]int, 0)
	str := getCompression(s, &stretchyStrCnt)

	var cnt int
	for _, word := range words {
		stretchyWordCnt := make([]int, 0)
		compressWord := getCompression(word, &stretchyWordCnt)
		if compressWord == str {
			var isStretchy = true
			for i, originStretchy := range stretchyStrCnt {
				queryStretchy := stretchyWordCnt[i]
				if queryStretchy == originStretchy {
					continue
				}

				if queryStretchy > originStretchy || originStretchy < 3 {
					isStretchy = false
					break
				}
			}

			if isStretchy {
				cnt++
			}
		}
	}

	return cnt
}

func getCompression(s string, letterCounter *[]int) string {
	builder := new(strings.Builder)
	builder.WriteByte(s[0])
	var counter = 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			counter++
		} else {
			if letterCounter != nil {
				*letterCounter = append(*letterCounter, counter)
			}

			builder.WriteByte(s[i])
			counter = 1
		}
	}

	if letterCounter != nil {
		*letterCounter = append(*letterCounter, counter)
	}
	return builder.String()
}
