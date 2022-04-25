package strobogrammatic_number

func FindStrobogrammatic(n int) []string {
	return helper(n, n)
}

func helper(n, m int) []string  {
	if n == 0 {
		return []string{""}
	}

	if n == 1 {
		return []string{"0", "1", "8"}
	}

	list := helper(n-2, m)
	res := make([]string, 0)

	for i := 0; i < len(list); i++ {
		s := list[i]
		if n != m {
			res = append(res, "0" + s +  "0")
		}
		res = append(res, "1" + s +  "1")
		res = append(res, "6" + s +  "9")
		res = append(res, "8" + s +  "8")
		res = append(res, "9" + s +  "6")
	}

	return res
}

const (
	zero = '0'
	one = '1'
	six = '6'
	eight = '8'
	nine = '9'
)

var (
	digits = []byte{zero, one, six, eight, nine}
	middleDigits = []byte{zero, one, eight}
	grammaticTable = map[byte]byte {
		zero: zero,
		one: one,
		six: nine,
		eight: eight,
		nine: six,
	}
)

func findStrobogrammatic(n int) []string {
	if n == 1 {
		return []string{string(zero), string(one), string(eight)}
	}

	eachStrobogrammatic := make([]byte, n)
	for i := 0; i < n; i++ {
		eachStrobogrammatic[i] = zero
	}

	strobogrammatic := make([]string, 0)
	even := n & 1 == 0
	generateStrobogrammatic(0, n/2, even, eachStrobogrammatic, &strobogrammatic)

	return strobogrammatic
}

func generateStrobogrammatic(k, n int, even bool, eachStrobogrammatic []byte, strobogrammatic *[]string) {
	if k >= n {
		return
	}

	var j int
	if k == 0 {
		j = 1
	}

	for ; j < len(digits); j++ {
		oldDigit := eachStrobogrammatic[k]
		eachStrobogrammatic[k] = digits[j]

		for i := 0; i < n; i++ {
			eachStrobogrammatic[len(eachStrobogrammatic)-i-1] = grammaticTable[eachStrobogrammatic[i]]
		}

		if !even {
			for _, mid := range middleDigits {
				eachStrobogrammatic[n] = mid

				if isStrobogrammaticAdded(strobogrammatic, 3, string(eachStrobogrammatic)) {
					break
				}

				*strobogrammatic = append(*strobogrammatic, string(eachStrobogrammatic))
			}
		} else if !isStrobogrammaticAdded(strobogrammatic, 1, string(eachStrobogrammatic)) {
			*strobogrammatic = append(*strobogrammatic, string(eachStrobogrammatic))
		}

		generateStrobogrammatic(k+1, n, even, eachStrobogrammatic, strobogrammatic)
		eachStrobogrammatic[k] = oldDigit
	}
}

func isStrobogrammaticAdded(strobogrammatic *[]string, i int, str string) bool {
	var added bool
	if len(*strobogrammatic) >= i && (*strobogrammatic)[len(*strobogrammatic)-i] == str {
		added = true
	}

	return added
}
