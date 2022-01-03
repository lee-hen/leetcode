package missing_ranges

import "strconv"

func FindMissingRanges(nums []int, lower int, upper int) []string {
	formatRange := func(lower int, upper int) string {
		if lower == upper {
			return strconv.Itoa(lower)
		}
		return strconv.Itoa(lower) + "->" + strconv.Itoa(upper)
	}

	result := make([]string, 0)
	prev := lower - 1
	for i := 0; i <= len(nums); i++ { // not handle nums[i] > upper
		var curr int
		if i < len(nums) {
			curr = nums[i]
		} else {
			curr = upper + 1
		}

		if prev + 1 <= curr - 1 {
			result = append(result, formatRange(prev + 1, curr - 1))
		}
		prev = curr
	}
	return result
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	missing := make([]string, 0)
	if len(nums) == 0 {
		appendMissing(&missing, lower, upper)
		return missing
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < lower {
			continue
		}

		var l, u int
		if nums[i] > upper {
			if i > 0 {
				l = nums[i-1]+1
			} else {
				l = lower
			}

			u = upper
			appendMissing(&missing, l, u)
			break
		}

		if i == 0 {
			if nums[0] > lower {
				l = lower
				u = nums[i]-1
				appendMissing(&missing, l, u)
			}
		}  else {
			l = nums[i-1]+1
			u = nums[i]-1
			appendMissing(&missing, l, u)
		}

		if i == len(nums)-1 {
			l = nums[i]+1
			u = upper
			appendMissing(&missing, l, u)
		}
	}

	return missing
}

func appendMissing(missing *[]string, lower, upper int) {
	if lower > upper {
		return
	}

	lowerUpperStr := strconv.Itoa(lower)
	if upper > lower {
		lowerUpperStr += "->"
		lowerUpperStr += strconv.Itoa(upper)
	}

	*missing = append(*missing, lowerUpperStr)
}
