package entities

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	BackgroundImg    *ebiten.Image
	ProjectileImg    *ebiten.Image
	EnemyImg         *ebiten.Image
	Enemies          []*Enemy
	Projectiles      []*Projectile
	Player           *Player
	SpawnTime        time.Time
	EnemiesDestroyed int
	gamepadIDsBuf    []ebiten.GamepadID
	gamepadIDs       map[ebiten.GamepadID]struct{}
}

const (
	ScreenHeight = 480
	ScreenWidth  = 640
)

var (
	startTime   = time.Now() // As far as i know this has to be here to happen only once
	displayTime = ""         // Keep it as a global variable so that we can display it after the game is over
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Define the game's screen size.
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if g.gamepadIDs == nil {
		g.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}

	// Manage connecting and disconnecting gamepad(s)
	g.gamepadIDsBuf = ebiten.AppendGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		g.gamepadIDs[id] = struct{}{}
	}
	for id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			delete(g.gamepadIDs, id)
		}
	}

	if gameOver {
		return nil
	}

	// Update Player
	g.Player.Update(g)

	// Update all Enemies
	for _, enm := range g.Enemies {
		enm.Update(g.Player)
	}

	// Update all Projectiles
	for _, p := range g.Projectiles {
		p.Update()
	}

	// Check for collisions between Projectiles and Enemies
	g.checkCollisions()

	// Check if it's time to spawn a new enemy
	if time.Since(g.SpawnTime).Seconds() >= 1 {
		g.spawnNewEnemy()

		// Reset the timer for next spawn
		g.SpawnTime = time.Now()
	}

	return nil
}

// The order of drawing elements on the screen determines their z-index
func (g *Game) Draw(screen *ebiten.Image) {

	// Do not proceed with logic unless the gamepad is connected
	if len(g.gamepadIDs) == 0 {
		ebitenutil.DebugPrint(screen, "Please connect your gamepad.")
		return
	}

	// Draw the background
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.BackgroundImg, opts)

	// Create string representing the amount of Enemies destroyed
	displayEnemiesDestroyed := strconv.Itoa(g.EnemiesDestroyed)
	stringToDisplay := fmt.Sprintln("Enemies destroyed: " + displayEnemiesDestroyed)

	// Create string representing elapsed time
	elapsedTime := time.Since(startTime)
	secondsPassed := int(math.Round(elapsedTime.Seconds()))
	displaySeconds := strconv.Itoa(secondsPassed % 60)
	minutesPassed := secondsPassed / 60
	displayMinutes := strconv.Itoa(minutesPassed % 60)
	hoursPassed := minutesPassed / 60
	displayHours := strconv.Itoa(hoursPassed % 24)
	if !gameOver {
		displayTime = displayHours + "h " + displayMinutes + "m " + displaySeconds + "s"
	}
	stringToDisplay += fmt.Sprintln("Elapsed time: " + displayTime)

	// Display amount of bolts user has
	stringToDisplay += fmt.Sprintln("Bolts available: " + strconv.Itoa(g.Player.BoltAmount))

	if gameOver {
		stringToDisplay += fmt.Sprintln("Game over! You've just got gobbled!")
	}

	// Display on the screen
	ebitenutil.DebugPrint(screen, stringToDisplay)

	if gameOver {
		return
	}

	// Draw all Enemies
	for _, enm := range g.Enemies {
		enm.Draw(screen)
	}

	// Draw all Projectiles
	for _, p := range g.Projectiles {
		p.Draw(screen)
	}

	// Draw the Player
	g.Player.Draw(screen)
}

// Custom functions with a Game receiver below

// Logic to spawn an enemy
func (g *Game) spawnNewEnemy() {
	// Create a new Enemy
	enm := &Enemy{
		img: g.EnemyImg,
	}

	// Randomly choose an edge (0=left, 1=top, 2=right, 3=bottom)
	edge := rand.Intn(4)
	switch edge {
	case 0:
		// Left edge
		enm.x = -float64(enm.img.Bounds().Dx()) * 2
		enm.y = rand.Float64() * ScreenHeight
	case 1:
		// Top edge
		enm.x = rand.Float64() * ScreenWidth
		enm.y = -float64(enm.img.Bounds().Dy()) * 2
	case 2:
		// Right edge
		enm.x = ScreenWidth
		enm.y = rand.Float64() * ScreenHeight
	case 3:
		// Bottom edge
		enm.x = rand.Float64() * ScreenWidth
		enm.y = ScreenHeight
	}

	// Add the new enemy to the Enemies slice
	g.Enemies = append(g.Enemies, enm)
}

func (g *Game) checkCollisions() {
	// Iterate over all Projectiles and Enemies to check for collisions
	for i := len(g.Projectiles) - 1; i >= 0; i-- {
		prj := g.Projectiles[i]
		for j := len(g.Enemies) - 1; j >= 0; j-- {
			enm := g.Enemies[j]
			if checkCollision(prj, enm) {
				// Remove the projectile and the enemy
				g.Projectiles = append(g.Projectiles[:i], g.Projectiles[i+1:]...)
				g.Enemies = append(g.Enemies[:j], g.Enemies[j+1:]...)
				g.EnemiesDestroyed++
				break
			}
		}
	}
}
