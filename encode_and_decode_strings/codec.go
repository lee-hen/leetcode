package encode_and_decode_strings

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {}

// Encode
// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	builder := new(strings.Builder)

	for _, s := range strs {
		chars := []byte(s)
		builder.WriteString(fmt.Sprintf("%d", chars))
		builder.WriteString(",")
	}

	return builder.String()
}

// Decode
// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	s := strings.Split(strs, ",")
	result := make([]string, 0)

	for _, c := range s {
		builder := new(strings.Builder)
		if len(c) > 0 {
			byteString := fmt.Sprintf("%s", c[1:len(c)-1])
			byteStrings := strings.Split(byteString, " ")
			for _, byteStr := range byteStrings {
				byteNumber, _ := strconv.ParseInt(byteStr, 10, 8)
				builder.WriteByte(byte(byteNumber))
			}

			if str := builder.String(); str != "\u0000" {
				result = append(result, str)
			} else {
				result = append(result, "")
			}
		}
	}

	return result
}
