package multiply_strings

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMultiply(t *testing.T) {
	num1, num2 := "2", "3"
	require.Equal(t, "6", Multiply(num1, num2))
	require.Equal(t, "6", multiply(num1, num2))

	num1, num2 = "123", "456"
	require.Equal(t, "56088", Multiply(num1, num2))
	require.Equal(t, "56088", multiply(num1, num2))

	num1, num2 = "999", "999"
	require.Equal(t, "998001", Multiply(num1, num2))
	require.Equal(t, "998001", multiply(num1, num2))
}
