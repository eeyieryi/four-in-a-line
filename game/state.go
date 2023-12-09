package game

type GameState struct {
	Board      *Board
	NextToPlay int
	BoardState BoardState
}

func NewGameState(board *Board, boardState BoardState, nextToPlay int) *GameState {
	if board == nil {
		newBoard := NewBoard(nil)
		board = &newBoard
	}
	return &GameState{
		Board:      board,
		BoardState: boardState,
		NextToPlay: nextToPlay,
	}
}
