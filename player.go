package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x        float64
	y        float64
	rotation float64
	image    *ebiten.Image // New field to store the loaded image
}

// In the Draw method of the Player struct
func (p *Player) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Get the dimensions of the image using Bounds()
	w, h := p.image.Bounds().Dx(), p.image.Bounds().Dy()

	// Translate to the center of the image before rotating
	opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate player
	// opts.GeoM.Rotate(p.rotation)

	// Mirror image
	currentAngle := math.Abs(math.Mod(p.rotation, 2*math.Pi))
	if currentAngle < math.Pi*0.5 || currentAngle > math.Pi*1.5 {
		opts.GeoM.Scale(-1, 1)
	}

	opts.GeoM.Translate(p.x, p.y)

	// Draw the player to the screen with the rotation options
	screen.DrawImage(p.image, opts)
}

// Shoot method of the Player struct
func (p *Player) Shoot(g *Game) {
	// Calculate the velocity based on the player's rotation
	vx := math.Cos(p.rotation) * projectileSpeed
	vy := math.Sin(p.rotation) * projectileSpeed

	// Create a new projectile with the calculated velocity
	proj := &Projectile{
		x:         p.x,
		y:         p.y,
		rotation:  p.rotation,
		velocityX: vx,
		velocityY: vy,
		img:       g.projectileImg,
	}

	g.projectiles = append(g.projectiles, proj)
}

// Update method of the Player struct
func (p *Player) Update() {
	// Update player's rotation based on user input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= rotationSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += rotationSpeed
	}

	// Clamp rotation angle to [0, 2π)
	p.rotation = math.Mod(p.rotation, 2*math.Pi)
}
