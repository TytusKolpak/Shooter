package main

import (
	"log"
	"shooter/entities"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
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

	// Load the Player image from the sprite sheet
	characterSheet, _, err := ebitenutil.NewImageFromFile(characterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// If importing from subdirectory they have to have first letter capitalized
	game := &entities.Game{
		SpawnTime:     time.Now(),
		BackgroundImg: backgroundImage,
		EnemyImg:      entities.LoadSpriteFromSheet(enemySheet, 2, 0),
		ProjectileImg: entities.LoadSpriteFromSheet(itemSheet, 0, 6),
		Player: &entities.Player{
			X:     entities.ScreenWidth / 2,
			Y:     entities.ScreenHeight / 2,
			Image: entities.LoadSpriteFromSheet(characterSheet, 2, 0),
		},
	}

	ebiten.SetWindowSize(entities.ScreenWidth, entities.ScreenHeight)
	ebiten.SetWindowTitle("Shoot them!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
