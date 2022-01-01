package plus_one

func PlusOne(digits []int) []int {
	// move along the input array starting from the end
	for idx := len(digits) - 1; idx >= 0; idx-- {
		// set all the nines at the end of array to zeros
		if digits[idx] == 9 {
			digits[idx] = 0
		} else { // here we have the rightmost not-nine
			// increase this rightmost not-nine by 1
			digits[idx]++
			// and the job is done
			return digits
		}
	}

	// we're here because all the digits are nines
	res := make([]int, len(digits)+1)
	res[0] = 1
	copy(res[1:], digits)
	return res
}

func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)

	var carry int
	for j := len(digits)-1; j >= 0; j-- {
		var plus int
		if j == len(digits)-1 {
			plus = digits[j] + 1 + carry
		} else {
			plus = digits[j] + carry
		}

		res[j+1] = plus % 10
		carry = plus / 10
	}
	if carry != 0 {
		res[0] = carry
	} else {
		return res[1:]
	}
	return res
}
