package main

import (
	"log"
	"shooter/entities"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// If importing from subdirectory they have to have first letter capitalized
	tileSheet, enemySheet, projectileSheet, playerSheet := entities.LoadSpriteSheets()

	game := &entities.Game{
		SpawnTime:     time.Now(),
		BackgroundImg: entities.GenerateBackground(tileSheet),
		EnemyImg:      entities.AddBoundingBox(entities.LoadSpriteFromSheet(enemySheet, 0, 2)),
		EnemySheet:    enemySheet,
		ProjectileImg: entities.AddBoundingBox(entities.LoadSpriteFromSheet(projectileSheet, 0, 6)),
		Player: &entities.Player{
			BoltAmount: entities.InitialBoltAmount,
			LoadTime:   time.Now(),
			X:          entities.ScreenWidth / 2,
			Y:          entities.ScreenHeight / 2,
			Img:        entities.AddBoundingBox(entities.LoadSpriteFromSheet(playerSheet, 4, 0)),
		},
		Stage: 1,
	}

	ebiten.SetWindowSize(entities.ScreenWidth, entities.ScreenHeight)
	ebiten.SetWindowTitle("Shoot them!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
