package longest_palindromic_substring

type substring struct {
	Left  int
	Right int
}
func (ss substring) Length() int {
	return ss.Right - ss.Left
}

func LongestPalindrome(str string) string {
	result := substring{0, 1}
	for i := 1; i < len(str); i++ {
		odd := GetLongestPalindromeFrom(str, i-1, i+1)
		even := GetLongestPalindromeFrom(str, i-1, i)
		longest := even
		if odd.Length() > even.Length() {
			longest = odd
		}
		if longest.Length() > result.Length() {
			result = longest
		}
	}
	return str[result.Left:result.Right]
}

func GetLongestPalindromeFrom(str string, leftIndex, rightIndex int) substring {
	for leftIndex >= 0 && rightIndex < len(str) {
		if str[leftIndex] != str[rightIndex] {
			break
		}
		leftIndex -= 1
		rightIndex += 1
	}
	return substring{leftIndex + 1, rightIndex}
}
