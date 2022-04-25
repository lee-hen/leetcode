package valid_parentheses

func IsValid(s string) bool {
	stack := make([]byte, 0)
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 {
				var top byte
				top, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if s[i] == ')' && top != '(' ||
					s[i] == ']' && top != '[' ||
					s[i] == '}' && top != '{' {
					return false
				}
			} else { // s = "]})"
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
