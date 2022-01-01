package jump_game

func CanJump(nums []int) bool {
	var traceIdx = len(nums)-1
	for j := len(nums)-2; j >= 0; j-- {
		if nums[j] != 0 && (nums[j] + j) >= traceIdx {
			traceIdx = j
		}
	}

	return traceIdx == 0
}

type Index int

const (
	GOOD Index = iota
	BAD
	UNKNOWN
)

func canJump(nums []int) bool {
	memo := make([]Index, len(nums))

	for i := 0; i < len(memo); i++ {
		memo[i] = UNKNOWN
	}
	memo[len(memo) - 1] = GOOD

	for i := len(nums)-2; i >= 0; i-- {
		furthestJump := Min(i+nums[i], len(nums)-1)
		for j := i + 1; j <= furthestJump; j++ {
			if memo[j] == GOOD {
				memo[i] = GOOD
				break
			}
		}
	}

	return canJumpFromPosition(0, nums) && memo[0] == GOOD
}

func canJumpFromPosition(position int, nums []int) bool {
	if position == len(nums)-1 {
		return true
	}

	furthestJump := Min(position+nums[position], len(nums)-1)
	for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
		if canJumpFromPosition(nextPosition, nums) {
			return true
		}
	}

	return false
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





