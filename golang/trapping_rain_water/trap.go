package trapping_rain_water

func Trap(height []int) int {
	left, right := 0, len(height) - 1
	var ans, leftMax, rightMax int

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}

func trap(height []int) int {
	var trap, currTrap int
	var leftMax = 0
	for i := range height {
		if height[i] < leftMax {
			currTrap += leftMax - height[i]
		}

		if leftMax <= height[i] {
			if leftMax != 0 {
				trap += currTrap
				currTrap = 0
			}
			leftMax = height[i]
		}
	}

	rightMax := 0
	for j := len(height)-1; leftMax > height[j]; j-- {
		if height[j] < rightMax {
			trap += rightMax - height[j]
		}

		if rightMax < height[j] {
			rightMax = height[j]
		}
	}

	return trap
}
