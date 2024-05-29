package entities

import (
	"time"
)

const (
	ScreenHeight = 480
	ScreenWidth  = 640

	CharacterSpriteSheetPath = "sprites/rogues.png"
	MonsterSpriteSheetPath   = "sprites/monsters.png"
	ItemSpriteSheetPath      = "sprites/items.png"
	TileSpriteSheetPath      = "sprites/tiles.png"

	PlayerSpeed       = 2
	InitialBoltAmount = 10
	rotationSpeed     = 0.03
	ProjectileSpeed   = 10
	StageDuration     = 30 // Seconds

	maxEnemySpeed = 0.75

	spriteSize = 32
)

var (
	startTime     = time.Now() // As far as i know this has to be here to happen only once
	pauseStart    = time.Now()
	pauseDuration time.Duration
	displayTime   = "" // Keep it as a global variable so that we can display it after the game is over
	gameOver      = false
	gamePaused    = false
)
