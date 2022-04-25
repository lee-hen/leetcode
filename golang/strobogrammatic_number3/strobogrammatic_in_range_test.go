package strobogrammatic_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 312498
	output := strobogrammaticInRange("1", "999999999999999")
	require.Equal(t, expected, output)
}

