package reverse_integer

import "math"

func Reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10


		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7)  {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8)  {
			return 0
		}


		rev = rev * 10 + pop
	}
	return rev
}


func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}

	y := x
	k := 1
	for y > 0 {
		y /= 10
		k *= 10
	}
	k /= 10

	z := x
	var r int
	for z > 0 {
		mod := z % 10
		r += k * mod
		k /= 10
		z /= 10
	}

	if r > math.MaxInt32 {
		return 0
	}

	return r
}
