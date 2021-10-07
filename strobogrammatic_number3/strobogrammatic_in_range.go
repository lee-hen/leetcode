package strobogrammatic_number

import "strconv"

func strobogrammaticInRange(low string, high string) int {
	l, _ := strconv.Atoi(low)
	h, _ := strconv.Atoi(high)

	var results = make(map[int]int)
	lowLen, highLen := len(low), len(high)

	for i := highLen; i >= lowLen; i-- {
		if _, ok := results[i]; ok {
			continue
		}

		helper(i, i, l, h, results)
	}

	var result int
	for _, v := range results {
		result += v
	}

	return result
}

func helper(n, m, l, h int, results map[int]int) []string  {
	if n == 0 {
		return []string{""}
	}

	if n == 1 {
		countNumberInRange(results, n, l, h,  0, 1, 8)
		return []string{"0", "1", "8"}
	}

	list := helper(n-2, m, l, h, results)

	res := make([]string, 0)
	for i := 0; i < len(list); i++ {
		s := list[i]
		if n != m {
			res = append(res, "0" + s +  "0")
		}

		s1 := "1" + s +  "1"
		s2 := "6" + s +  "9"
		s3 := "8" + s +  "8"
		s4 := "9" + s +  "6"

		number1, _ := strconv.Atoi(s1)
		number2, _ := strconv.Atoi(s2)
		number3, _ := strconv.Atoi(s3)
		number4, _ := strconv.Atoi(s4)

		countNumberInRange(results, len(s)+2, l, h, number1, number2, number3, number4)

		res = append(res, []string{s1, s2, s3, s4}...)
	}

	return res
}

func countNumberInRange(results map[int]int, numOfDigit, l, h int, numbers ...int) {
	for _, number := range numbers {
		if number >= l && number <= h {
			results[numOfDigit]++
		}
	}
}
