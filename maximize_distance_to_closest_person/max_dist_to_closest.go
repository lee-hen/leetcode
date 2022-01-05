package maximize_distance_to_closest_person

func MaxDistToClosest(seats []int) int {
	k := 0 // current longest group of empty seats
	ans := 0

	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			k = 0
		} else {
			k++
			ans = Max(ans, (k + 1) / 2)
		}
	}

	for i := 0; i < len(seats); i++  {
		if seats[i] == 1 {
			ans = Max(ans, i)
			break
		}
	}

	for i := len(seats)-1; i >= 0; i-- {
		 if seats[i] == 1 {
			 ans = Max(ans, len(seats)-1-i)
			 break
		 }
	}

	return ans
}

func maxDistToClosestTwoPointer(seats []int) int {
	prev := -1
	future := 0
	ans := 0

	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			prev = i
		} else {
			for future < len(seats) && seats[future] == 0 || future < i {
				future++
			}
			var left int
			if prev == -1 {
				left = len(seats)
			} else {
				left = i - prev
			}

			var right int
			if future == len(seats) {
				right = len(seats)
			} else {
				right = future - i
			}

			ans = Max(ans, Min(left, right))
		}
	}

	return ans
}

func maxDistToClosest(seats []int) int {
	distances := make([]int, len(seats))
	var idx = -1
	for i := range seats {
		if seats[i] == 1 {
			idx = i
		} else if idx != -1 {
			distances[i] = i-idx
		}
	}

	idx = len(seats)
	for j := len(seats)-1; j >= 0; j-- {
		if seats[j] == 1 {
			idx = j
		}  else if idx != len(seats) {
			if distances[j] > idx-j || distances[j] == 0  {
				distances[j] = idx-j
			}
		}
	}

	return Max(distances[0], distances...)
}

func Max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}
