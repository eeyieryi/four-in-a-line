package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func handleInput(g *Game) error {
	// if ebiten.IsKeyPressed(ebiten.KeyQ) {
	// 	return ErrTerminated
	// }

	switch g.currentScene {
	case PlayingScene:
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			if g.selectedColumn-1 > 0 {
				g.selectedColumn -= 1
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			if g.selectedColumn+1 <= COLUMNS_MAX {
				g.selectedColumn += 1
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			valid := IsValidMove(*g.State.Board, g.State.NextToPlay, g.selectedColumn)
			if !valid {
				log.Printf("Not a valid move! Player (%d) at Column (%d)", g.State.NextToPlay, g.selectedColumn)
			} else {
				newBoard := AddPiece(*g.State.Board, g.State.NextToPlay, g.selectedColumn)
				newState, winningPieces := GetBoardState(newBoard)

				g.State.Board = &newBoard
				g.State.BoardState = newState

				switch g.State.BoardState {
				case PlayerOneWinState, PlayerTwoWinState, DrawState:
					g.winningPieces = winningPieces
					g.currentScene = GameOverScene
				default:
					switch g.State.NextToPlay {
					case PlayerOne:
						g.State.NextToPlay = PlayerTwo
					case PlayerTwo:
						g.State.NextToPlay = PlayerOne
					}
				}
			}
		}
	case GameOverScene:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.StartGame()
		}
	}

	return nil
}
