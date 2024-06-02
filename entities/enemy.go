package entities

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	img   *ebiten.Image
	reach float64
	x, y  float64
}

func (enm *Enemy) Update(p *Player) {
	// Calculate the difference in position
	dx := (p.X - spriteSize/2) - enm.x
	dy := (p.Y - spriteSize/2) - enm.y

	// Calculate the distance to the destination point (enemy to player)
	distance := math.Sqrt(dx*dx + dy*dy)

	// If the image is very close to the center, stop moving
	if distance < enm.reach {
		gameOver = true
		return
	}

	// Normalize the direction vector (up to 1)
	dirX := dx / distance
	dirY := dy / distance

	// Move the image towards the center
	enm.x += dirX * maxEnemySpeed
	enm.y += dirY * maxEnemySpeed
}

func (enm *Enemy) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Set the position of the image
	opts.GeoM.Translate(enm.x, enm.y)

	// Draw the image to the screen with the scaling options
	screen.DrawImage(enm.img, opts)
}
