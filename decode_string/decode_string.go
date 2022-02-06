package decode_string

import (
	"container/list"
	"math"
	"strings"
	"unicode"
)

var idx = 0

func DecodeString(s string) string {
	result := new(strings.Builder)
	for idx < len(s) && s[idx] != ']' {
		if !unicode.IsDigit(rune(s[idx])) {
			result.WriteByte(s[idx])
			idx++
		} else {
			k := 0
			// build k while next character is a digit
			for idx < len(s) && unicode.IsDigit(rune(s[idx])) {
				k = k * 10 + int(s[idx] - '0')
				idx++
			}

			// ignore the opening bracket '['
			idx++
			decodedString := DecodeString(s)
			// ignore the closing bracket ']'
			idx++
			// build k[decodedString] and append to the result
			for k > 0 {
				result.WriteString(decodedString)
				k--
			}
		}
	}
	return result.String()
}

func decodeStringTwoStack(s string) string {
	countStack := list.New() // int
	stringStack := list.New() // *strings.Builder
	currentString := new(strings.Builder)

	k := 0
	for i := range s {
		ch := s[i]
		if unicode.IsDigit(rune(ch)) {
			k = k * 10 + int(ch - '0')
		} else if ch == '[' {
			// push the number k to countStack
			countStack.PushBack(k)
			// push the currentString to stringStack
			stringStack.PushBack(currentString)
			// reset currentString and k
			currentString = new(strings.Builder)
			k = 0
		} else if ch == ']' {
			back := stringStack.Back()
			stringStack.Remove(back)
			decodedString := back.Value.(*strings.Builder)

			back = countStack.Back()
			countStack.Remove(back)
			currentK := back.Value.(int)

			// decode currentK[currentString] by appending currentString k times
			for  currentK > 0  {
				decodedString.WriteString(currentString.String())
				currentK--
			}

			currentString = decodedString
		} else {
			currentString.WriteByte(ch)
		}
	}
	return currentString.String()
}

func decodeString(s string) string {
	chars := make([]byte, 0)

	for i := range s {
		c := s[i]
		if c == ']' {
			partial := make([]byte, 0)
			for len(chars) > 0 {
				var char byte
				char, chars = chars[len(chars)-1], chars[:len(chars)-1]
				if char == '[' {
					break
				}
				partial = append(partial, char)
			}

			var k int
			var digit float64
			for len(chars) > 0 {
				char := chars[len(chars)-1]
				if !unicode.IsDigit(rune(char)) {
					break
				}

				chars = chars[:len(chars)-1]
				k += int(char - '0') * int(math.Pow(10, digit))
				digit++
			}

			for k > 0 {
				for j := len(partial)-1; j >= 0; j-- {
					chars = append(chars, partial[j])
				}
				k--
			}
		}  else {
			chars = append(chars, c)
		}
	}

	return string(chars)
}
