package main

import (
	"log"
	"shooter/entities"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	var err error

	// Load the background image (put these in entities somewhere)
	tileSheet, _, err := ebitenutil.NewImageFromFile(entities.TileSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the item image from the sprite sheet
	enemySheet, _, err := ebitenutil.NewImageFromFile(entities.MonsterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the item image from the sprite sheet
	itemSheet, _, err := ebitenutil.NewImageFromFile(entities.ItemSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the Player image from the sprite sheet
	characterSheet, _, err := ebitenutil.NewImageFromFile(entities.CharacterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// If importing from subdirectory they have to have first letter capitalized
	game := &entities.Game{
		SpawnTime:     time.Now(),
		BackgroundImg: entities.GenerateBackground(tileSheet),
		EnemyImg:      entities.LoadSpriteFromSheet(enemySheet, 2, 0),
		ProjectileImg: entities.LoadSpriteFromSheet(itemSheet, 0, 6),
		Player: &entities.Player{
			BoltAmount: entities.InitialBoltAmount,
			LoadTime:   time.Now(),
			X:          entities.ScreenWidth / 2,
			Y:          entities.ScreenHeight / 2,
			Image:      entities.LoadSpriteFromSheet(characterSheet, 4, 0),
		},
	}

	ebiten.SetWindowSize(entities.ScreenWidth, entities.ScreenHeight)
	ebiten.SetWindowTitle("Shoot them!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
