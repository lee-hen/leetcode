package swap_adjacent_in_lr_string

func CanTransform(start string, end string) bool {
	rightIndices := make([]int, 0)
	leftIndices := make([]int, 0)

	for i := range start {
		if start[i] == 'R' {
			rightIndices = append(rightIndices, i)
		}
		if start[i] == 'L' {
			leftIndices = append(leftIndices, i)
		}
	}

	leftIndices = ReverseSlice(leftIndices)
	rightIndices = ReverseSlice(rightIndices)

	for i := range end {
		c := end[i]

		var left = -1
		if len(leftIndices) > 0 {
			left = leftIndices[len(leftIndices)-1]
		}

		var right = -1
		if len(rightIndices) > 0 {
			right = rightIndices[len(rightIndices)-1]
		}

		if c == 'L' {
			if left == -1 {
				return false
			}

			leftIndices = leftIndices[:len(leftIndices)-1]

			if left < i {
				return false
			}
			// start RL end LR
			// start RXXLX end LXRXX
			if right != -1 && left != -1 && right < left {
				return false
			}
		}

		if c == 'R' {
			if right == -1 {
				return false
			}

			rightIndices = rightIndices[:len(rightIndices)-1]
			if right > i {
				return false
			}
		}
	}

	if len(leftIndices) > 0 || len(rightIndices) > 0 {
		return false
	}

	return true
}

func ReverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
