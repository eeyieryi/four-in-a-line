package game

import (
	"reflect"
)

type Board [][]int

const (
	ROWS_MAX    = 6
	COLUMNS_MAX = 7
)

const (
	PlayerOne = 1
	PlayerTwo = 2
)

type BoardState int

type Piece struct {
	row int
	col int
}

const (
	DrawState         BoardState = -1
	OngoingState      BoardState = 0
	PlayerOneWinState BoardState = 1
	PlayerTwoWinState BoardState = 2
)

func emptyBoard() Board {
	var rows Board
	for i := 0; i < ROWS_MAX; i++ {
		row := make([]int, COLUMNS_MAX)
		rows = append(rows, row)
	}
	return rows
}

func copyBoard(other Board) Board {
	rows := emptyBoard()

	for i := 0; i < ROWS_MAX; i++ {
		for j := 0; j < COLUMNS_MAX; j++ {
			rows[i][j] = other[i][j]
		}
	}

	return rows
}

func NewBoard(state Board) Board {
	board := emptyBoard()

	if state != nil {
		board = copyBoard(state)
	}

	return board
}

func AddPiece(board Board, player int, column int) Board {
	newState := NewBoard(board)
	for i := ROWS_MAX - 1; i >= 0; i-- {
		if newState[i][column-1] == 0 {
			newState[i][column-1] = player
			break
		}
	}
	return newState
}

func IsValidMove(state Board, player int, column int) bool {
	newState := AddPiece(NewBoard(state), player, column)
	return !reflect.DeepEqual(state, newState)
}

func GetBoardState(board Board) (BoardState, []Piece) {
	found := false

	winner := 0
	winningPieces := make([]Piece, 4)

	for i := ROWS_MAX - 1; i >= 0; i-- {
		if found {
			break
		}
		for j := 0; j < COLUMNS_MAX; j++ {
			value := board[i][j]
			if value == 0 {
				continue
			} else {
				if i-3 >= 0 {
					if board[i-1][j] == value &&
						board[i-2][j] == value &&
						board[i-3][j] == value {

						a := Piece{i, j}
						b := Piece{i - 1, j}
						c := Piece{i - 2, j}
						d := Piece{i - 3, j}
						winningPieces = []Piece{a, b, c, d}

						winner = value
						found = true
						break
					}
					if j+3 < COLUMNS_MAX {
						if board[i-1][j+1] == value &&
							board[i-2][j+2] == value &&
							board[i-3][j+3] == value {

							a := Piece{i, j}
							b := Piece{i - 1, j + 1}
							c := Piece{i - 2, j + 2}
							d := Piece{i - 3, j + 3}
							winningPieces = []Piece{a, b, c, d}

							winner = value
							found = true
							break
						}
					}
					if j-3 >= 0 {
						if board[i-1][j-1] == value &&
							board[i-2][j-2] == value &&
							board[i-3][j-3] == value {

							a := Piece{i, j}
							b := Piece{i - 1, j - 1}
							c := Piece{i - 2, j - 2}
							d := Piece{i - 3, j - 3}
							winningPieces = []Piece{a, b, c, d}

							winner = value
							found = true
							break
						}
					}
				}
				if j+3 < COLUMNS_MAX {
					if board[i][j+1] == value &&
						board[i][j+2] == value &&
						board[i][j+3] == value {

						a := Piece{i, j}
						b := Piece{i, j + 1}
						c := Piece{i, j + 2}
						d := Piece{i, j + 3}
						winningPieces = []Piece{a, b, c, d}

						winner = value
						found = true
						break
					}
				}
			}
		}
	}

	if !found {
		emptyCell := false

		for i := 0; i < ROWS_MAX; i++ {
			if emptyCell {
				break
			}
			for j := 0; j < COLUMNS_MAX; j++ {
				if board[i][j] == 0 {
					emptyCell = true
					break
				}
			}
		}

		if !emptyCell {
			return -1, nil
		}
	}

	return BoardState(winner), winningPieces
}
