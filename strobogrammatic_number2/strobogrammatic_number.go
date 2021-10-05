package strobogrammatic_number

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

func FindStrobogrammatic(n int) []string {
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
