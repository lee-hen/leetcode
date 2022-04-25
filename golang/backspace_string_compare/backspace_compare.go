package backspace_string_compare

func BackspaceCompare(s string, t string) bool {
	backspacing := func(s string) string {
		buffer := make([]byte, 0)
		for i := range s {
			if s[i] != '#' {
				buffer = append(buffer, s[i])
			} else if len(buffer) > 0 {
				buffer = buffer[:len(buffer)-1]
			}
		}

		return string(buffer)
	}

	return backspacing(s) == backspacing(t)
}
