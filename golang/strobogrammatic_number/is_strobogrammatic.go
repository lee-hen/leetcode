package strobogrammatic_number

const (
	zero = '0'
	one = '1'
	six = '6'
	eight = '8'
	nine = '9'
)

var (
	grammaticTable = map[byte]byte {
		zero: zero,
		one: one,
		six: nine,
		eight: eight,
		nine: six,
	}
)

func IsStrobogrammatic(num string) bool {
	var i int
	for j := len(num)-1; i < j; i, j = i+1, j-1 {
		if x, ok := grammaticTable[num[i]]; !ok || x != num[j] {
			return false
		}
	}

	if len(num) & 1 != 0 {
		if x, ok := grammaticTable[num[i]]; !ok || x == six || x == nine {
			return false
		}
	}
	return true
}
