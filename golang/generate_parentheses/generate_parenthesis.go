package generate_parentheses

import "strings"

const (
	open = '('
	closed = ')'
)

func GenerateParenthesis(n int) []string {
	return backtrack(0, 0, n,  new(strings.Builder))
}

func backtrack(left, right, max int, prefix *strings.Builder) []string {
	res := make([]string, 0)
	if prefix.Len() == max * 2 {
		res = append(res, prefix.String())
		return res
	}

	if left < max {
		prefix.WriteByte(open)
		res = append(res, backtrack(left+1, right, max, prefix)...)

		cpy := prefix.String()
		prefix.Reset()
		prefix.WriteString(cpy[:len(cpy)-1])
	}

	if right < left {
		prefix.WriteByte(closed)
		res = append(res, backtrack(left, right+1, max, prefix)...)

		cpy := prefix.String()
		prefix.Reset()
		prefix.WriteString(cpy[:len(cpy)-1])
	}
	return res
}

func generateParenthesis(n int) []string {
	prefix := make([]byte, 2 * n)
	return helper(n, n, 0, prefix)
}

func helper(leftRem, rightRem, idx int, prefix []byte) []string {
	res := make([]string, 0)
	if leftRem < 0 || rightRem < leftRem{
		return res
	}

	if leftRem == 0 && rightRem == 0 {
		res = append(res, string(prefix))
	} else {
		prefix[idx] = open
		res = append(res, helper(leftRem-1, rightRem, idx+1, prefix)...)
		prefix[idx] = closed
		res = append(res, helper(leftRem, rightRem-1, idx+1, prefix)...)
	}
	return res
}

func otherSolution(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
	} else {
		for c := 0; c < n; c++ {
			for _, left := range otherSolution(c) {
				for _, right := range otherSolution(n-1-c) {
					res = append(res, string(open) + left + string(closed) + right)
				}
			}
		}
	}
	return res
}
