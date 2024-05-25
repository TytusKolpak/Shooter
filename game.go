package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var startTime = time.Now() // As far as i know this has to be here to happen only once

type Game struct {
	backgroundImg    *ebiten.Image
	projectileImg    *ebiten.Image
	enemyImg         *ebiten.Image
	enemies          []*Enemy
	projectiles      []*Projectile
	player           *Player
	spawnTime        time.Time
	enemiesDestroyed int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Define the game's screen size.
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	// Update player
	g.player.Update()

	// Update all enemies
	for _, enm := range g.enemies {
		enm.Update()
	}

	// Update all projectiles
	for _, p := range g.projectiles {
		p.Update()
	}

	// Check for collisions between projectiles and enemies
	g.checkCollisions()

	// Check if it's time to spawn a new enemy
	if time.Since(g.spawnTime).Seconds() >= 1 {
		g.spawnNewEnemy()

		// Reset the timer for next spawn
		g.spawnTime = time.Now()
	}

	// Handle player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.player.Shoot(g)
	}

	return nil
}

func (g *Game) spawnNewEnemy() {
	// Create a new Enemy
	enm := &Enemy{
		img: g.enemyImg,
	}

	// Randomly choose an edge (0=left, 1=top, 2=right, 3=bottom)
	edge := rand.Intn(4)
	switch edge {
	case 0:
		// Left edge
		enm.x = -float64(enm.img.Bounds().Dx()) * 2
		enm.y = rand.Float64() * screenHeight
	case 1:
		// Top edge
		enm.x = rand.Float64() * screenWidth
		enm.y = -float64(enm.img.Bounds().Dy()) * 2
	case 2:
		// Right edge
		enm.x = screenWidth
		enm.y = rand.Float64() * screenHeight
	case 3:
		// Bottom edge
		enm.x = rand.Float64() * screenWidth
		enm.y = screenHeight
	}

	// Add the new enemy to the enemies slice
	g.enemies = append(g.enemies, enm)
}

func (g *Game) checkCollisions() {
	// Iterate over all projectiles and enemies to check for collisions
	for i := len(g.projectiles) - 1; i >= 0; i-- {
		prj := g.projectiles[i]
		for j := len(g.enemies) - 1; j >= 0; j-- {
			enm := g.enemies[j]
			if checkCollision(prj, enm) {
				// Remove the projectile and the enemy
				g.projectiles = append(g.projectiles[:i], g.projectiles[i+1:]...)
				g.enemies = append(g.enemies[:j], g.enemies[j+1:]...)
				g.enemiesDestroyed++
				break
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the background
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.backgroundImg, opts)

	// Draw all enemies
	for _, enm := range g.enemies {
		enm.Draw(screen)
	}

	// Draw all projectiles
	for _, p := range g.projectiles {
		p.Draw(screen)
	}

	// Draw the player
	g.player.Draw(screen)

	// Create string representing the amount of enemies destroyed
	displayEnemiesDestroyed := strconv.Itoa(g.enemiesDestroyed)
	stringToDisplay := "Enemies destroyed: " + displayEnemiesDestroyed + ". "

	// Create string representing elapsed time
	elapsedTime := time.Since(startTime)
	secondsPassed := int(math.Round(elapsedTime.Seconds()))
	displaySeconds := strconv.Itoa(secondsPassed % 60)
	minutesPassed := secondsPassed / 60
	displayMinutes := strconv.Itoa(minutesPassed % 60)
	hoursPassed := minutesPassed / 60
	displayHours := strconv.Itoa(hoursPassed % 24)
	displayTime := displayHours + "h " + displayMinutes + "m " + displaySeconds + "s"
	stringToDisplay += "Elapsed time: " + displayTime

	// Display on the screen
	ebitenutil.DebugPrint(screen, stringToDisplay)
}
