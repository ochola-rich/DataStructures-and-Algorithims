package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	row := 0
	board := [8]int{}
	for i := range board {
		board[i] = -1
	}
	backtrack(row, &board)
}

func backtrack(row int, board *[8]int) {
	if row == len(board) {
		for _, col := range board {
			z01.PrintRune(rune(col + '0' + 1))
		}
		z01.PrintRune('\n')
		return
	}
	for col := 0; col < len(board); col++ {
		if Safe(board, row, col) {
			board[row] = col
			backtrack(row+1, board)
			board[row] = -1
		}
	}
}

func Safe(board *[8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]+i == col+row || board[i]-i == col-row {
			return false
		}
	}
	return true
}
