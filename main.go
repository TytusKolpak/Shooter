package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenHeight             = 480
	screenWidth              = 640
	characterSpriteSheetPath = "sprites/rogues.png"
	monsterSpriteSheetPath   = "sprites/monsters.png"
	itemSpriteSheetPath      = "sprites/items.png"
	backgroundImagePath      = "sprites/background.png"
)

func main() {
	var err error

	// Load the background image
	backgroundImage, _, err := ebitenutil.NewImageFromFile(backgroundImagePath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the item image from the sprite sheet
	enemySheet, _, err := ebitenutil.NewImageFromFile(monsterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the item image from the sprite sheet
	itemSheet, _, err := ebitenutil.NewImageFromFile(itemSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the player image from the sprite sheet
	characterSheet, _, err := ebitenutil.NewImageFromFile(characterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		spawnTime:     time.Now(),
		backgroundImg: backgroundImage,
		enemyImg:      loadSpriteFromSheet(enemySheet, 2, 0),
		projectileImg: loadSpriteFromSheet(itemSheet, 0, 6),
		player: &Player{
			x:     screenWidth / 2,
			y:     screenHeight / 2,
			image: loadSpriteFromSheet(characterSheet, 2, 0),
		},
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Shoot them!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
