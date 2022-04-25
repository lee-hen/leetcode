package next_closest_time

import (
	"fmt"
	"strings"
)

func NextClosestTime(time string) string {
	closest := []byte(time)
	minimumClosest := "24:00"
	for i := len(time)-1; i >= 0; i-- {
		if time[i] == ':' {
			continue
		}

		for j := range time {
			if time[j] == ':' {
				continue
			}

			if time[j] > time[i] {
				closest[i] = time[j]
				if string(closest) > time && string(closest[:2]) <= "23" && string(closest[3:]) <= "59" {
					minimumClosest = Min(minimumClosest, string(closest))
				}
				closest[i] = time[i]
			}
		}
	}

	var minimumDigit = Min("9", strings.Split(time, "")...)
	closest = []byte(minimumClosest)
	if minimumClosest == "24:00" {
		return fmt.Sprintf("%s%s:%s%s", minimumDigit, minimumDigit, minimumDigit, minimumDigit)
	}

	for i := range minimumClosest {
		if minimumClosest[i] > time[i] {
			for j := i+1; j < len(minimumClosest); j++ {
				if minimumClosest[j] != ':' {
					closest[j] = minimumDigit[0]
				}
			}
			break
		}
	}

	return string(closest)
}

func Min(arg string, rest ...string) string {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
