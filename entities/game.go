package entities

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	TileSheet        *ebiten.Image
	BackgroundImg    *ebiten.Image
	ProjectileImg    *ebiten.Image
	EnemySheet       *ebiten.Image
	EnemyImg         *ebiten.Image
	Enemies          []*Enemy
	Projectiles      []*Projectile
	Player           *Player
	SpawnTime        time.Time
	EnemiesDestroyed int
	Stage            int
	gamepadIDsBuf    []ebiten.GamepadID
	gamepadIDs       map[ebiten.GamepadID]struct{}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Define the game's screen size.
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	// --------------------------- Game behavior ---------------------------
	if g.gamepadIDs == nil {
		g.gamepadIDs = map[ebiten.GamepadID]struct{}{}
	}

	// Log the gamepad connection events.
	g.gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDsBuf[:0])
	for _, id := range g.gamepadIDsBuf {
		g.gamepadIDs[id] = struct{}{}
	}
	for id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			delete(g.gamepadIDs, id)
		}
	}

	if gameOver {
		if len(g.gamepadIDs) != 0 {
			for id := range g.gamepadIDs {
				// Center left button (so start or select or options or menu)
				if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton6) {
					// Create "new" Game instance, by resetting variables
					g.ResetGame()
				}
				// Center right button
				if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton7) {
					os.Exit(0)
				}
			}
		} else {
			if ebiten.IsKeyPressed(ebiten.KeyR) {
				g.ResetGame()
			}
			if ebiten.IsKeyPressed(ebiten.KeyQ) {
				os.Exit(0)
			}
		}
		return nil
	} else {
		if len(g.gamepadIDs) != 0 {
			for id := range g.gamepadIDs {
				// "Is...JustPressed", so that it will not fire multiple times
				if inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton6) {
					gamePaused = !gamePaused
					pauseStart = time.Now().Add(-pauseDuration)
				}
			}
		} else {
			if ebiten.IsKeyPressed(ebiten.KeyP) {
				gamePaused = !gamePaused
				pauseStart = time.Now().Add(-pauseDuration)
			}
		}
	}

	if gamePaused {
		return nil
	}

	// --------------------------- In game objects behavior ---------------------------
	// Update Player
	g.Player.Update(g)

	// Update all Enemies
	for _, enm := range g.Enemies {
		enm.Update(g.Player)
	}

	// Update all Projectiles
	for i := len(g.Projectiles) - 1; i >= 0; i-- {
		prj := g.Projectiles[i]
		prjx, prjy := prj.x, prj.y

		// If projectile is out of the screen - remove it
		if prjx < 0 || prjx > ScreenWidth || prjy < 0 || prjy > ScreenHeight {
			g.Projectiles = append(g.Projectiles[:i], g.Projectiles[i+1:]...)
		}
		prj.Update()
	}
	// Check for collisions between Projectiles and Enemies
	g.checkCollisions()

	// Check for projectile pickup by the Player
	g.checkPickups()

	// Check if it's time to spawn a new enemy and not last stage
	if g.Stage != 4 && time.Since(g.SpawnTime).Seconds() >= SpawnInterval {
		g.spawnNewEnemy()

		// Reset the timer for next spawn
		g.SpawnTime = time.Now()
	}

	g.controlGameStage()

	return nil
}

// The order of drawing elements on the screen determines their z-index
func (g *Game) Draw(screen *ebiten.Image) {
	// It does have to be redrawn despite being static since ebiten clears screen every frame
	screen.DrawImage(g.BackgroundImg, nil)

	// Do not proceed with logic unless the gamepad is connected
	if len(g.gamepadIDs) == 0 {
		ebitenutil.DebugPrintAt(screen, "Using Keyboard", 0, ScreenHeight-15)
	} else {
		ebitenutil.DebugPrintAt(screen, "Using Gamepad", 0, ScreenHeight-15)
	}

	// Create string representing the amount of Enemies destroyed
	displayEnemiesDestroyed := strconv.Itoa(g.EnemiesDestroyed)
	stringToDisplay := fmt.Sprintln("Enemies destroyed: " + displayEnemiesDestroyed)

	// Create string representing elapsed time
	elapsedTime := time.Since(startTime)
	if gamePaused {
		pauseDuration = time.Since(pauseStart)
		ebitenutil.DebugPrintAt(screen, "Game Paused", ScreenWidth/2-40, ScreenHeight/2-10)
	}
	elapsedTime -= pauseDuration
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

	// Display on the screen
	ebitenutil.DebugPrint(screen, stringToDisplay)

	if gameOver {
		restartText := ""
		quitText := ""
		if len(g.gamepadIDs) != 0 {
			restartText = "Left Center button to Restart"
			quitText = "Right Center button to Quit"
		} else {
			restartText = "R key to Restart"
			quitText = "Q key to Quit"
		}

		messageFrameX := ScreenWidth/2 - 90
		messageFrameY := ScreenHeight/2 - 50
		lineHeight := 15

		ebitenutil.DebugPrintAt(screen, restartText, messageFrameX, messageFrameY)
		ebitenutil.DebugPrintAt(screen, quitText, messageFrameX, messageFrameY+lineHeight)

		if g.Stage == 4 && len(g.Enemies) == 0 {
			ebitenutil.DebugPrintAt(screen, "YOU WIN! :D", messageFrameX, messageFrameY+2*lineHeight)
		} else {
			ebitenutil.DebugPrintAt(screen, "Game over! You've just got gobbled!", messageFrameX, messageFrameY+2*lineHeight)
		}
		return
	}

	// Draw all Projectiles
	for _, p := range g.Projectiles {
		p.Draw(screen)
	}

	// Draw all Enemies
	for _, enm := range g.Enemies {
		enm.Draw(screen)
	}

	// Draw the Player
	g.Player.Draw(screen)
}

// Custom functions with a Game receiver below

func (g *Game) ResetGame() {
	g.Enemies = nil
	g.Projectiles = nil
	g.EnemyImg = AddBoundingBox(LoadSpriteFromSheet(g.EnemySheet, 0, 2))
	g.Stage = 1

	g.EnemiesDestroyed = 0
	g.Player.BoltAmount = 10
	g.Player.X = ScreenWidth / 2
	g.Player.Y = ScreenHeight / 2
	g.Player.BoltAmount = InitialBoltAmount

	startTime = time.Now()
	displayTime = ""
	pauseDuration = 0
	gameOver = false
}

// Logic to spawn an enemy
func (g *Game) spawnNewEnemy() {
	// Create a new Enemy
	enm := &Enemy{
		img: g.EnemyImg,
	}

	if g.Stage == 1 {
		enm.reach = 0.8 * spriteSize
	} else if g.Stage == 2 {
		enm.reach = spriteSize
	} else {
		enm.reach = 1.2 * spriteSize
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

		// If the projectile is laying on the ground then don't kill enemies with it
		if !prj.active {
			break
		}

		for j := len(g.Enemies) - 1; j >= 0; j-- {
			enm := g.Enemies[j]
			if checkCollision(prj, enm) {
				// Remove the enemy
				g.Enemies = append(g.Enemies[:j], g.Enemies[j+1:]...)
				g.EnemiesDestroyed++

				// Leave the projectile
				prj.velocityX, prj.velocityY = 0, 0
				prj.active = false
				break
			}
		}
	}
}

func (g *Game) checkPickups() {
	for i := len(g.Projectiles) - 1; i >= 0; i-- {
		prj := g.Projectiles[i]

		// Don't consider it for pickup if it's flying
		if prj.active {
			continue
		}

		if checkPickup(prj, g.Player) {
			g.Player.addBolt()
			// Remove that projectile
			g.Projectiles = append(g.Projectiles[:i], g.Projectiles[i+1:]...)
		}
	}
}

func (g *Game) controlGameStage() {
	// If 1 minute has passed
	if g.Stage == 1 && time.Since(startTime).Seconds() > StageDuration {
		g.Stage = 2
		g.EnemyImg = AddBoundingBox(LoadSpriteFromSheet(g.EnemySheet, 0, 0))
	} else if g.Stage == 2 && time.Since(startTime).Seconds() > 2*StageDuration {
		g.Stage = 3
		g.EnemyImg = AddBoundingBox(LoadSpriteFromSheet(g.EnemySheet, 0, 1))
	} else if g.Stage == 3 && time.Since(startTime).Seconds() > 3*StageDuration {
		g.Stage = 4
	} else if g.Stage == 4 && len(g.Enemies) == 0 {
		// You win
		gameOver = true
	}
}
