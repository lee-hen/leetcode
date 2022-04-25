package tic_tac_toe

import "math"

type TicTacToe struct {
	rows []int
	cols []int
	diagonal int
	antiDiagonal int
}

func New(n int) *TicTacToe {
	t := new(TicTacToe)
	t.rows = make([]int, n)
	t.cols = make([]int, n)

	return t
}

func (t *TicTacToe) Move(row, col int, player int) int {
	currentPlayer := -1
	if player == 1 {
		currentPlayer = 1
	}

	t.rows[row] += currentPlayer
	t.cols[col] += currentPlayer

	// update diagonal
	if row == col {
		t.diagonal += currentPlayer
	}

	//update anti diagonal
	if col == len(t.cols) - row - 1 {
		t.antiDiagonal += currentPlayer
	}

	n := len(t.rows)
	// check if the current player wins
	if int(math.Abs(float64(t.rows[row]))) == n ||
		int(math.Abs(float64(t.cols[col]))) == n ||
		int(math.Abs(float64(t.diagonal))) == n ||
		int(math.Abs(float64(t.antiDiagonal))) == n {

		return player
	}
	// No one wins
	return 0
}

type Piece int

const (
	Empty Piece = iota
	Red
	Blue
)

type ticTacToe struct {
	board [][]Piece
	player int
	turn Piece
}

func constructor(n int) ticTacToe {
	t := new(ticTacToe)
	t.board = make([][]Piece, n)
	for i := range t.board {
		t.board[i] = make([]Piece, n)
	}

	return *t
}

func (t *ticTacToe) Move(row int, col int, player int) int {
	if row >= len(t.board) || col >= len(t.board) || row < 0 || col < 0 {
		return 0
	}

	if t.turn == Empty {
		t.turn = Red
	} else if t.player != player {
		if t.turn == Red {
			t.turn = Blue
		} else {
			t.turn = Red
		}
	}

	t.player = player
	t.board[row][col] = t.turn

	if HasWon(row, col, t.board) {
		return player
	}

	return 0
}

func HasWon(row, col int, board [][]Piece) bool {
	var hasWinner = false

	if board[row][0] != Empty && !hasWinner {
		/* Check Rows */
		for i := 1; i < len(board[row]); i++ {
			if board[row][i] != board[row][i-1] {
				hasWinner = false
				break
			} else {
				hasWinner = true
			}
		}
	}

	if board[0][col] != Empty && !hasWinner {
		/* Check Columns */
		for i := 1; i < len(board); i++ {
			if board[i][col] != board[i-1][col] {
				hasWinner = false
				break
			} else {
				hasWinner = true
			}
		}
	}

	if board[0][0] != Empty && !hasWinner {
		/* Check Diagonal */
		for i, j := 1, 1; i < len(board) && j < len(board[len(board)-1]); i, j = i+1, j+1 {
			if board[i][j] != board[i-1][j-1] {
				hasWinner = false
				break
			} else {
				hasWinner = true
			}
		}
	}

	if board[0][len(board)-1] != Empty && !hasWinner {
		for i, j := 1, len(board)-2;  i < len(board) && j >= 0; i, j = i+1, j-1 {
			if board[i][j] != board[i-1][j+1]  {
				hasWinner = false
				break
			} else {
				hasWinner = true
			}
		}
	}

	return hasWinner
}
