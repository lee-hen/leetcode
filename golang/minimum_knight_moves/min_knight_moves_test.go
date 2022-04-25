package minimum_knight_moves

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase0(t *testing.T) {
	require.Equal(t, 4, MinKnightMovesDFS(2, -2))
	require.Equal(t, 4, MinKnightMovesBFS(2, -2))
	require.Equal(t, 4, minKnightMovesDFS(2, -2))
	require.Equal(t, 4, minKnightMovesBFS(2, -2))
}

func TestCase1(t *testing.T) {
	require.Equal(t, 4, MinKnightMovesDFS(5, 5))
	require.Equal(t, 4, MinKnightMovesBFS(5, 5))
	require.Equal(t, 4, minKnightMovesDFS(5, 5))
	require.Equal(t, 4, minKnightMovesBFS(5, 5))
}

func TestCase2(t *testing.T) {
	require.Equal(t, 1, MinKnightMovesDFS(2, 1))
	require.Equal(t, 1, MinKnightMovesBFS(2, 1))
	require.Equal(t, 1, minKnightMovesDFS(2, 1))
	require.Equal(t, 1, minKnightMovesBFS(2, 1))
}

func TestCase3(t *testing.T) {
	require.Equal(t, 6, MinKnightMovesDFS(2, 12))
	require.Equal(t, 6, MinKnightMovesBFS(2, 12))
	require.Equal(t, 6, minKnightMovesDFS(2, 12))
	require.Equal(t, 6, minKnightMovesBFS(2, 12))
}

func TestCase4(t *testing.T) {
	require.Equal(t, 56, MinKnightMovesDFS(2, 112))
	require.Equal(t, 56, MinKnightMovesBFS(2, 112))
	require.Equal(t, 56, minKnightMovesDFS(2, 112))
	require.Equal(t, 56, minKnightMovesBFS(2, 112))
}

func TestCase5(t *testing.T) {
	require.Equal(t, 15, MinKnightMovesDFS(21, 20))
	require.Equal(t, 15, MinKnightMovesBFS(21, 20))
	require.Equal(t, 15, minKnightMovesDFS(21, 20))
	require.Equal(t, 15, minKnightMovesBFS(21, 20))
}

func TestCase7(t *testing.T) {
	require.Equal(t, 15, MinKnightMovesDFS(-20, 21))
	require.Equal(t, 15, MinKnightMovesBFS(-20, 21))
	require.Equal(t, 15, minKnightMovesDFS(-20, 21))
	require.Equal(t, 15, minKnightMovesBFS(-20, 21))
}

func TestCase6(t *testing.T) {
	require.Equal(t, 15, MinKnightMovesDFS(20, -21))
	require.Equal(t, 15, MinKnightMovesBFS(20, -21))
	require.Equal(t, 15, minKnightMovesDFS(20, -21))
	require.Equal(t, 15, minKnightMovesBFS(20, -21))
}

func TestCase8(t *testing.T) {
	require.Equal(t, 135, MinKnightMovesDFS(21, 270))
	require.Equal(t, 135, MinKnightMovesBFS(21, 270))
	require.Equal(t, 135, minKnightMovesDFS(21, 270))
	require.Equal(t, 135, minKnightMovesBFS(21, 270))
}

func TestCase9(t *testing.T) {
	require.Equal(t, 135, MinKnightMovesDFS(-270, 21))
	require.Equal(t, 135, MinKnightMovesBFS(-270, 21))
	require.Equal(t, 135, minKnightMovesDFS(-270, 21))
	require.Equal(t, 135, minKnightMovesBFS(-270, 21))
}

func TestCase10(t *testing.T) {
	require.Equal(t, 135, MinKnightMovesDFS(270, -21))
	require.Equal(t, 135, MinKnightMovesBFS(270, -21))
	require.Equal(t, 135, minKnightMovesDFS(270, -21))
	require.Equal(t, 135, minKnightMovesBFS(270, -21))
}

func TestCase11(t *testing.T) {
	require.Equal(t, 74, MinKnightMovesDFS(62, 148))
	require.Equal(t, 74, MinKnightMovesBFS(62, 148))
	require.Equal(t, 74, minKnightMovesDFS(62, 148))
	require.Equal(t, 74, minKnightMovesBFS(62, 148))
}

func TestCase12(t *testing.T) {
	require.Equal(t, 78, MinKnightMovesDFS(-34, -156))
	require.Equal(t, 78, MinKnightMovesBFS(-34, -156))
	require.Equal(t, 78, minKnightMovesDFS(-34, -156))
	require.Equal(t, 78, minKnightMovesBFS(-34, -156))
}

func TestCase13(t *testing.T) {
	require.Equal(t, 79, MinKnightMovesDFS(68, -157))
	require.Equal(t, 79, MinKnightMovesBFS(68, -157))
	require.Equal(t, 79, minKnightMovesDFS(68, -157))
	require.Equal(t, 79, minKnightMovesBFS(68, -157))
}

func TestCase14(t *testing.T) {
	require.Equal(t, 79, MinKnightMovesDFS(-68, 157))
	require.Equal(t, 79, MinKnightMovesBFS(-68, 157))
	require.Equal(t, 79, minKnightMovesDFS(-68, 157))
	require.Equal(t, 79, minKnightMovesBFS(-68, 157))
}

func TestCase15(t *testing.T) {
	require.Equal(t, 2, MinKnightMovesDFS(1, 1))
	require.Equal(t, 2, MinKnightMovesBFS(1, 1))
	require.Equal(t, 2, minKnightMovesDFS(1, 1))
	require.Equal(t, 2, minKnightMovesBFS(1, 1))
}

func TestCase16(t *testing.T) {
	require.Equal(t, 3, MinKnightMovesDFS(0, 1))
	require.Equal(t, 3, MinKnightMovesBFS(0, 1))
	require.Equal(t, 3, minKnightMovesDFS(0, 1))
	require.Equal(t, 3, minKnightMovesBFS(0, 1))
}
