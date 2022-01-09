package find_first_and_last_position_of_element_in_sorted_array

func SearchRange(nums []int, target int) []int {
	//result := helper(nums, target, 0, len(nums)-1, &Range{-1, -1})
	result := helper(nums, target, 0, len(nums)-1)
	return []int{result.left, result.right}
}

type Range struct {
	left, right int
}

func helper(a []int, x, low, high int) *Range {
	r := &Range{-1, -1}
	if low > high {
		return r
	}

	mid := (low + high) / 2
	if a[mid] == x {
		rL := helper(a, x, low, mid-1)
		if rL.left == -1 {
			rL.left = mid
		}

		rR := helper(a, x, mid+1, high)
		if rR.right == -1 {
			rR.right = mid
		}
		return &Range{rL.left, rR.right}
	} else if a[mid] < x {
		return helper(a, x, mid + 1, high)
	} else {
		return helper(a, x, low, mid - 1)
	}
}

func helper1(a []int, x, low, high int, r *Range) *Range {
	if low > high {
		return r
	}

	mid := (low + high) / 2
	if a[mid] == x {
		r = helper1(a, x, low, mid-1, r)
		if r.left == -1 || r.left > mid {
			r.left = mid
		}

		r = helper1(a, x, mid+1, high, r)
		if r.right == -1 || r.right < mid {
			r.right = mid
		}

		return r
	} else if a[mid] < x {
		return helper1(a, x, mid + 1, high, r)
	} else {
		return helper1(a, x, low, mid - 1, r)
	}
}
