package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"

	"github.com/eeyieryi/four-in-a-row/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	colorRed = color.RGBA{
		R: 0xFF,
		G: 0x00,
		B: 0x00,
		A: 0xFF,
	}
	colorYellow = color.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0x00,
		A: 0xFF,
	}
	colorBlue = color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0xFF,
		A: 0xFF,
	}
)

var (
	ErrTerminated = errors.New("terminated")
)

const (
	SCREEN_WIDTH  = 610
	SCREEN_HEIGHT = 340
)

type Game struct {
	title string

	gameImg *ebiten.Image

	board board.Board

	selectedColumn int // From 1 - 7
	currentPlayer  int // 1 for PlayerOne or 2 for PlayerTwo
	currentState   int // -1 for draw, 0 for continue, 1 for PlayerOne Win, 2 for PlayerTwo Win
}

func (g *Game) Setup() {
	g.title = "four-in-a-row"
	g.board = board.NewBoard(nil)
	g.gameImg = ebiten.NewImage(board.COLUMNS_MAX*20, board.ROWS_MAX*20+25)
	g.selectedColumn = 1
	g.currentPlayer = 1
	g.currentState = 0
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ErrTerminated
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if g.selectedColumn-1 > 0 {
			g.selectedColumn -= 1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if g.selectedColumn+1 <= board.COLUMNS_MAX {
			g.selectedColumn += 1
		}
	}

	switch g.currentState {
	case 0:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			valid := board.IsValidMove(g.board, g.currentPlayer, g.selectedColumn)
			if !valid {
				log.Printf("Not a valid move! Player: %d at Column: %d", g.currentPlayer, g.selectedColumn)
			} else {
				newBoard := board.AddPiece(g.board, g.currentPlayer, g.selectedColumn)
				newState := board.GetBoardState(newBoard)

				g.board = board.NewBoard(newBoard)
				g.currentState = newState

				switch g.currentPlayer {
				case 1:
					g.currentPlayer = 2
				case 2:
					g.currentPlayer = 1
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	{
		// TODO: Temporary
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Selected Column: %d", g.selectedColumn), screen.Bounds().Max.X-120, 5)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Current Player: %d", g.currentPlayer), screen.Bounds().Max.X-120, 25)
		ebitenutil.DebugPrintAt(screen, "Player 1", screen.Bounds().Max.X-120, 60)
		vector.DrawFilledCircle(screen, float32(screen.Bounds().Max.X-60), 67.5, 5, colorRed, true)
		ebitenutil.DebugPrintAt(screen, "Player 2", screen.Bounds().Max.X-120, 80)
		vector.DrawFilledCircle(screen, float32(screen.Bounds().Max.X-60), 87.5, 5, colorYellow, true)
	}

	{
		g.gameImg.Clear()
		g.gameImg.Fill(colorBlue)
		vector.DrawFilledRect(g.gameImg, 0, 0, board.COLUMNS_MAX*20, 20+5, color.Black, false)
	}

	{
		var circleColor color.Color
		if g.currentState != 0 {
			circleColor = color.White
		} else {
			switch g.currentPlayer {
			case 1:
				circleColor = colorRed
			case 2:
				circleColor = colorYellow
			}
		}
		vector.DrawFilledCircle(g.gameImg, float32(g.selectedColumn*20-20+10), 15.0, 5.0, circleColor, true)
	}

	{
		for i := 0; i < board.ROWS_MAX; i++ {
			for j := 0; j < board.COLUMNS_MAX; j++ {
				var myColor color.Color
				switch g.board[i][j] {
				case 0:
					myColor = color.White
				case 1:
					myColor = colorRed
				case 2:
					myColor = colorYellow
				}
				vector.DrawFilledCircle(g.gameImg, float32(j*20+10), float32(i*20+20+15), 5, myColor, true)
			}
		}
	}

	{
		geoM := ebiten.GeoM{}
		geoM.Translate(20, 10)
		screen.DrawImage(g.gameImg, &ebiten.DrawImageOptions{
			GeoM: geoM,
		})
	}

	{
		var theMessage string
		switch g.currentState {
		case -1:
			theMessage = "DRAW!"
		case 1:
			theMessage = "Player 1 WIN!"
		case 2:
			theMessage = "Player 2 WIN!"
		}
		// TODO: Temporary
		ebitenutil.DebugPrintAt(screen, theMessage, 20, 5)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH / 2, SCREEN_HEIGHT / 2
}

func main() {
	myGame := Game{}
	myGame.Setup()
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle(myGame.title)
	if err := ebiten.RunGame(&myGame); err != nil {
		if err == ErrTerminated {
			return
		}
		log.Fatal(err)
	}
}
