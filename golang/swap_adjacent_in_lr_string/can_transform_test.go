package swap_adjacent_in_lr_string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanTransform(t *testing.T) {
	start := "RXXLRXRXL"
	end := "XRLXXRRLX"

	require.True(t, CanTransform(start, end))
}

func TestCanTransform1(t *testing.T) {
	start := "XLXRRXXLRX"
	end := "XLXRRXXLXR"

	require.True(t, CanTransform(start, end))
}

func TestCanTransform3(t *testing.T) {
	start := "XXXXXXRXXLXRXXXXXRXXXXXRXXXXXLXXXLXLXXRXXXXXLXXXXX"
	end := "XXRXXXXLXXRXXXRXXXXRXXXXXLXXLXXXXXXLXXXXRXXXXLXXXX"

	require.False(t, CanTransform(start, end))
}

func TestCanTransform4(t *testing.T) {
	start := "XLLR"
	end := "LXLX"

	require.False(t, CanTransform(start, end))
}

func TestCanTransform5(t *testing.T) {
	start := "RL"
	end := "LR"

	require.False(t, CanTransform(start, end))
}

func TestCanTransform6(t *testing.T) {
	start := "RLX"
	end := "XLR"

	require.False(t, CanTransform(start, end))
}

func TestCanTransform7(t *testing.T) {
	start := "XXRXL"
	end := "LXRXX"

	require.False(t, CanTransform(start, end))
}


func TestCanTransform8(t *testing.T) {
	start := "X"
	end := "L"

	require.False(t, CanTransform(start, end))
}
