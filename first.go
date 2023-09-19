package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	_ "image/png"
)

type firstGame struct {
	player *ebiten.Image
	xloc   int
	yloc   int
	score  int
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("First Class Example")
	playerPict, _, err := ebitenutil.NewImageFromFile("galleon.png")
	if err != nil {
		fmt.Println("Error loading image:", err)
	}
	ourGame := firstGame{player: playerPict,
		xloc: 500,
		yloc: 500}
	err = ebiten.RunGame(&ourGame)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}

func (game *firstGame) Update() error {
	game.xloc = game.xloc + 5
	if game.xloc > 1000 {
		game.xloc = -72
	}
	return nil
}

func (game *firstGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Cornflowerblue)
	drawOps := ebiten.DrawImageOptions{}
	drawOps.GeoM.Reset()
	drawOps.GeoM.Translate(float64(game.xloc), float64(game.yloc))
	screen.DrawImage(game.player, &drawOps)
}

func (game firstGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight //by default, just return the current dimensions
}
