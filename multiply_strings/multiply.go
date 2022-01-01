package multiply_strings

import (
	"strconv"
	"strings"
)

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	firstNumber := ReverseBytes([]byte(num1))
	secondNumber := ReverseBytes([]byte(num2))

	// To store the multiplication result of each digit of secondNumber with firstNumber.
	n := len(firstNumber) + len(secondNumber)
	answer := make([]byte, n)
	for i := range answer {
		answer[i] = '0'
	}

	for place2 := 0; place2 < len(secondNumber); place2++ {
		digit2 := secondNumber[place2] - '0'

		// For each digit in secondNumber multiply the digit by all digits in firstNumber.
		for place1 := 0; place1 < len(firstNumber); place1++ {
			digit1 := firstNumber[place1] - '0'

			// The number of zeros from multiplying to digits depends on the
			// place of digit2 in secondNumber and the place of the digit1 in firstNumber.
			currentPos := place1 + place2

			// The digit currently at position currentPos in the answer string
			// is carried over and summed with the current result.
			carry := answer[currentPos] - '0'
			multiplication := digit1*digit2 + carry

			// Set the ones place of the multiplication result.
			answer[currentPos] = multiplication%10 + '0'

			// Carry the tens place of the multiplication result by
			// adding it to the next position in the answer array.
			value := (answer[currentPos+1] - '0') + multiplication/10
			answer[currentPos+1] = value + '0'
		}
	}

	// Pop excess 0 from the rear of answer.
	if answer[len(answer)-1] == '0' {
		answer = answer[:len(answer)-1]
	}

	return string(ReverseBytes(answer))
}

func ReverseBytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	matrix := make([][]int, len(num2))
	for i := range num2 {
		var carryDown int
		for j := range num1 {
			digit := int((num1[j]-'0')*(num2[i]-'0'))
			matrix[i] = append(matrix[i], digit/10 + carryDown)
			if i > 0 {
				matrix[i][len(matrix[i])-1] += matrix[i-1][len(matrix[i])]
			}
			carryDown = digit%10
		}

		matrix[i] = append(matrix[i], carryDown)
	}

	buf := make([]int, len(matrix) + len(matrix[0])-1)
	var carryUp int
	k := len(buf)-1
	for i, j := len(matrix)-1, len(matrix[len(matrix)-1])-1; j >= 0; j-- {
		digit := matrix[i][j]%10 + carryUp
		buf[k] = digit%10
		carryUp = matrix[i][j]/10 + digit/10
		k--

		if j == 0 {
			for i = i-1; i >= 0; i-- {
				digit = matrix[i][0]%10 + carryUp
				buf[k] = digit % 10
				carryUp = matrix[i][0]/10 + digit/10
				k--
			}
		}
	}

	if buf[0] == 0 {
		buf = buf[1:]
	}

	res := new(strings.Builder)
	for j := range buf {
		res.WriteString(strconv.Itoa(buf[j]))
	}

	return res.String()
}
