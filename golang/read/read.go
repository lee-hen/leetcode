package read

var Read = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	buffer := make([]byte, 0)
	buf4 := make([]byte, 4)
	for length := read4(buf4); length != 0; length = read4(buf4) {
		buffer = append(buffer, buf4[:length]...)
	}

	ptr := 0
	return func(buf []byte, n int) int {
		currLength := 0
		for i := ptr; currLength < n && i < len(buffer); i++ {
			buf[currLength] = buffer[i]
			currLength++
		}
		ptr += currLength
		return currLength
	}
}
