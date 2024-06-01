// Package entities deals with the character and behavior of all game elements
// including Player, Enemy, Projectile, but also things like global parameters and Game itself.
package entities

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	x, y                 float64 // These address the CENTER of an image
	velocityX, velocityY float64
	rotation             float64
	active               bool
	img                  *ebiten.Image
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Shrink the projectile
	// opts.GeoM.Scale(0.5, 0.5)

	// Set the position of the projectile for reference to rotation before any rotation
	opts.GeoM.Translate(-16, -16)

	// Rotate the projectile
	opts.GeoM.Rotate(p.rotation + math.Pi*0.75) // Rotating extra from the original angle

	// Set the position of the projectile
	opts.GeoM.Translate(p.x, p.y)

	// Draw the projectile to the screen
	screen.DrawImage(p.img, opts)
}

func (p *Projectile) Update() {
	// Move the projectile based on its velocity
	p.x += p.velocityX
	p.y += p.velocityY
}

// These 2 check functions might be connected using interfaces
func checkCollision(p *Projectile, e *Enemy) bool {
	// This seems to hit earlier from top and from left than from right and bottom

	// Calculate the difference in position
	dx := p.x - e.x
	dy := p.y - e.y

	// Calculate the distance to the destination point (enemy to player)
	distance := math.Sqrt(dx*dx + dy*dy)

	// If the distance is smaller than the sprite size, then it reaches
	isReaching := distance < spriteSize

	return isReaching
}

func checkPickup(prj *Projectile, plr *Player) bool {
	// Calculate the difference in position
	dx := prj.x - plr.X
	dy := prj.y - plr.Y

	// Calculate the distance to the destination point (enemy to player)
	distance := math.Sqrt(dx*dx + dy*dy)

	// If the distance is smaller than the sprite size, then it reaches
	isReaching := distance < spriteSize // This will be changed to player.reach

	return isReaching
}
