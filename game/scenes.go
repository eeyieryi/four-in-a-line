package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Scene int

const (
	PlayingScene Scene = iota
	GameOverScene
)

var (
	sceneImg  *ebiten.Image
	playerImg *ebiten.Image
	boardImg  *ebiten.Image
)

const (
	pieceRadius   float32 = 5.0
	pieceDiameter         = pieceRadius * 2
	tileSize      float32 = 20
	offsetX               = tileSize
	offsetY               = tileSize * 2
	boardWidth            = COLUMNS_MAX * tileSize
	boardHeight           = ROWS_MAX * tileSize
)

var (
	sceneDrawingOptions ebiten.DrawImageOptions
)

func setupScenes() {
	sceneImg = ebiten.NewImage(ScreenWidth, ScreenHeight)

	const width = int(boardWidth + offsetY)
	playerImg = ebiten.NewImage(width, int(offsetY))
	boardImg = ebiten.NewImage(width, int(boardHeight+tileSize+offsetY))

	sceneGeoM := ebiten.GeoM{}
	sceneGeoM.Scale(2, 2)
	sceneDrawingOptions = ebiten.DrawImageOptions{GeoM: sceneGeoM}
}

func drawScenes(g *Game, screen *ebiten.Image) {
	screen.Clear()
	sceneImg.Clear()
	boardImg.Clear()

	switch g.currentScene {
	case PlayingScene:
		playerImg.Clear()
		drawPlayer(g, playerImg)
		sceneImg.DrawImage(playerImg, &ebiten.DrawImageOptions{})

		drawGame(g, screen)
	case GameOverScene:
		drawGame(g, screen)
	}
}

func drawGame(g *Game, screen *ebiten.Image) {
	drawInfo(g, screen)
	drawBoard(g, boardImg)
	sceneImg.DrawImage(boardImg, &ebiten.DrawImageOptions{})
	screen.DrawImage(sceneImg, &sceneDrawingOptions)
}

func drawInfo(g *Game, img *ebiten.Image) {
	var msgX int = img.Bounds().Dx()/2 + int(offsetX*3)
	var initialMsgY int = int(offsetY)
	var gap int = int(tileSize)
	var msgOffsetY int = gap

	{
		msgY := initialMsgY + msgOffsetY
		msg := fmt.Sprintf("Selected Column: %d", g.selectedColumn)
		text.Draw(img, msg, MediumFont, msgX, msgY, color.White)
	}

	{
		msgOffsetY += gap * 2
		msgY := initialMsgY + msgOffsetY
		msg := fmt.Sprintf("Current Player: %d", g.State.NextToPlay)
		text.Draw(img, msg, MediumFont, msgX, msgY, color.White)
	}

	{
		var playerCircleX float32 = float32(msgX) + 120
		var playerCircleY float32
		const playerCircleRadius = pieceDiameter

		msgOffsetY += gap

		{
			msgOffsetY += gap
			msgY := initialMsgY + msgOffsetY

			msg := "Player 1"
			text.Draw(img, msg, MediumFont, msgX, msgY, color.White)

			playerCircleY = float32(msgY) - playerCircleRadius/2
			vector.DrawFilledCircle(img, playerCircleX, playerCircleY, playerCircleRadius, ColorRed, true)
		}

		msgOffsetY += gap / 2

		{
			msgOffsetY += gap
			msgY := initialMsgY + msgOffsetY

			msg := "Player 2"
			text.Draw(img, msg, MediumFont, msgX, msgY, color.White)

			playerCircleY = float32(msgY) - playerCircleRadius/2
			vector.DrawFilledCircle(img, playerCircleX, playerCircleY, playerCircleRadius, ColorYellow, true)
		}
	}

	if g.currentScene == GameOverScene {
		{
			var msg string
			var msgColor color.Color = color.White
			switch g.State.BoardState {
			case DrawState:
				msg = "DRAW!"
				msgColor = color.White
			case PlayerOneWinState:
				msg = "Player 1 WIN!"
				msgColor = ColorRed
			case PlayerTwoWinState:
				msg = "Player 2 WIN!"
				msgColor = ColorYellow
			}
			text.Draw(img, msg, LargeFont, int(offsetX*2), int(offsetY+tileSize), msgColor)
		}

		{
			msgOffsetY += gap * 5
			msgY := initialMsgY + msgOffsetY

			msg := "TO START A NEW GAME\n    PRESS SPACE"
			text.Draw(img, msg, MediumFont, msgX-gap, msgY, color.White)
		}
	}
}

func drawPlayer(g *Game, img *ebiten.Image) {
	var circleColor color.Color
	if g.State.BoardState != OngoingState {
		circleColor = color.White
	} else {
		switch g.State.NextToPlay {
		case PlayerOne:
			circleColor = ColorRed
		case PlayerTwo:
			circleColor = ColorYellow
		}
	}
	selectedColumn := float32(g.selectedColumn)
	var cx float32 = selectedColumn*tileSize - pieceDiameter + offsetX
	var cy float32 = offsetY - pieceDiameter
	vector.DrawFilledCircle(img, cx, cy, pieceRadius, circleColor, true)
}

func drawBoard(g *Game, img *ebiten.Image) {
	vector.DrawFilledRect(img, offsetX, offsetY, boardWidth, boardHeight, ColorBlue, false)

	{
		board := *g.State.Board
		for i := 0; i < ROWS_MAX; i++ {
			for j := 0; j < COLUMNS_MAX; j++ {
				var tileColor color.Color
				switch board[i][j] {
				case 0:
					tileColor = color.White
				case 1:
					tileColor = ColorRed
				case 2:
					tileColor = ColorYellow
				}
				row := float32(i)
				col := float32(j)
				cx := col*tileSize + pieceDiameter + offsetX
				cy := row*tileSize + pieceDiameter + offsetY
				vector.DrawFilledCircle(img, cx, cy, pieceRadius, tileColor, true)
			}
		}
	}

	{
		switch g.State.BoardState {
		case PlayerOneWinState, PlayerTwoWinState:
			if g.winningPieces != nil {
				fromPiece := g.winningPieces[0]
				toPiece := g.winningPieces[3]

				fromX := float32(fromPiece.col)*tileSize + offsetX + pieceDiameter
				fromY := float32(fromPiece.row)*tileSize + offsetY + pieceDiameter
				toX := float32(toPiece.col)*tileSize + offsetX + pieceDiameter
				toY := float32(toPiece.row)*tileSize + offsetY + pieceDiameter

				vector.StrokeLine(img, fromX, fromY, toX, toY, 1, color.White, false)
			}
		}
	}
}
