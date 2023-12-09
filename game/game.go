package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 360
)

type Game struct {
	Title string
	State *GameState

	currentScene   Scene
	selectedColumn int // From 1 to 7
	winningPieces  []Piece
}

func (g *Game) StartGame() {
	whoStartsGame := PlayerOne

	if g.State != nil {
		switch g.State.BoardState {
		case PlayerOneWinState:
			whoStartsGame = PlayerTwo
		case PlayerTwoWinState:
			whoStartsGame = PlayerOne
		}
	}

	g.State = NewGameState(nil, OngoingState, whoStartsGame)
	g.selectedColumn = 1
	g.currentScene = PlayingScene
}

func (g *Game) Setup() {
	g.Title = "Four in a row"
	setupFonts()
	setupScenes()
	g.StartGame()
}

func (g *Game) Update() (err error) {
	err = handleInput(g)
	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawScenes(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth int, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
