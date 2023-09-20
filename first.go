package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	_ "image/png"
	"math/rand"
)

type firstGame struct {
	player    *ebiten.Image
	xloc      int
	yloc      int
	score     int
	treasures []coinPile
}

type coinPile struct {
	picture *ebiten.Image
	xloc    int
	yloc    int
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("First Class Example")
	playerPict, _, err := ebitenutil.NewImageFromFile("galleon.png")
	if err != nil {
		fmt.Println("Error loading image:", err)
	}
	coinPict, _, err := ebitenutil.NewImageFromFile("coins.png")
	if err != nil {
		fmt.Println("Errorloading image:", err)
	}
	allTreasures := make([]coinPile, 0, 15)
	for pileNum := 0; pileNum < 10; pileNum++ {
		allTreasures = append(allTreasures,
			newCoinPile(928, 928, coinPict))
	}
	ourGame := firstGame{player: playerPict,
		xloc:      500,
		yloc:      500,
		treasures: allTreasures,
	}
	err = ebiten.RunGame(&ourGame)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}

func newCoinPile(maxWidth int, maxHeight int, pict *ebiten.Image) coinPile {
	return coinPile{
		picture: pict,
		xloc:    rand.Intn(maxWidth),
		yloc:    rand.Intn(maxHeight),
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
	for shipNum := 0; shipNum < 3; shipNum++ {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(game.xloc-shipNum*80),
			float64(game.yloc))
		screen.DrawImage(game.player, &drawOps)
	}
	for _, pile := range game.treasures {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(float64(pile.xloc), float64(pile.yloc))
		screen.DrawImage(pile.picture, &drawOps)
	}
}

func (game firstGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight //by default, just return the current dimensions
}
