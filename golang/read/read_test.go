package read

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct{
		query int
		expected int
	}{
		{
			1, 1,
		},
		{
			2, 2,
		},
		{
			1, 0,
		},
	}

	fp, file := 0, "abc"


	/**
	 * The read4 API is already defined for you.
	 *
	 *     read4 := func(buf4 []byte) int
	 *
	 * // Below is an example of how the read4 API can be called.
	 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
	 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
	 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
	 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
	 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
	 */
	read4 := func(buf4 []byte) int {
		if fp >= len(file) {
			return 0
		}

		length := 0
		for i := fp; i < fp + 4 && i < len(file); i++ {
			buf4[length] = file[i]
			length++
		}

		fp += length
		return length
	}

	read := Read(read4)
	for _, test := range tests {
		buf := make([]byte, len(file))
		require.Equal(t, test.expected, read(buf, test.query))
	}
}
