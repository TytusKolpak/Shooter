package entities

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PlayerSpeed       = 1
	InitialBoltAmount = 10
)

type Player struct {
	LoadTime       time.Time
	BoltAmount     int
	BoltShotBefore bool
	X              float64
	Y              float64
	Rotation       float64
	Image          *ebiten.Image // New field to store the loaded image
}

// In the Draw method of the Player struct
func (p *Player) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Get the dimensions of the image using Bounds()
	w, h := p.Image.Bounds().Dx(), p.Image.Bounds().Dy()

	// Translate to the center of the image before rotating
	opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Mirror image
	currentAngle := math.Abs(math.Mod(p.Rotation, 2*math.Pi))
	if currentAngle < math.Pi*0.5 || currentAngle > math.Pi*1.5 {
		opts.GeoM.Scale(-1, 1)
	}

	opts.GeoM.Translate(p.X, p.Y)

	// Draw the Player to the screen with the rotation options
	screen.DrawImage(p.Image, opts)
}

// Shoot method of the Player struct
func (p *Player) Shoot(g *Game) {
	// Calculate the velocity based on the Player's rotation
	vx := math.Cos(p.Rotation) * Projectilespeed
	vy := math.Sin(p.Rotation) * Projectilespeed

	// Create a new projectile with the calculated velocity
	proj := &Projectile{
		x:         p.X,
		y:         p.Y,
		rotation:  p.Rotation,
		velocityX: vx,
		velocityY: vy,
		img:       g.ProjectileImg,
	}

	g.Projectiles = append(g.Projectiles, proj)
}

// Update method of the Player struct
func (p *Player) Update(g *Game) {
	// Update Player position and the angle he shoots at it has to be in if statement
	// because user more than likely will use more than one at the same time

	// Arrow Keys control angle to shoot at
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.Rotation -= rotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.Rotation += rotationSpeed
	}

	// WSAD keys control movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= PlayerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += PlayerSpeed
	}

	// Handle Player shooting (enable one shot at a time and only if player has bolts available)
	BoltToBeShotNow := ebiten.IsKeyPressed(ebiten.KeySpace)
	if BoltToBeShotNow && !p.BoltShotBefore && p.BoltAmount > 0 {
		p.BoltShotBefore = true
		g.Player.Shoot(g)
		p.removeBolt()
	}
	p.BoltShotBefore = BoltToBeShotNow

	// Clamp rotation angle to [0, 2π)
	p.Rotation = math.Mod(p.Rotation, 2*math.Pi)

	// every 0.5 s add a bolt to the player
	if time.Since(p.LoadTime).Seconds() >= 0.5 {
		p.addBolt()

		// Reset the timer for next spawn
		p.LoadTime = time.Now()
	}
}

func (p *Player) addBolt() {
	if p.BoltAmount < 20 {
		p.BoltAmount += 1
	}
}

func (p *Player) removeBolt() {
	p.BoltAmount -= 1
}
