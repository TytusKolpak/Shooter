package entities

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	img  *ebiten.Image
	x, y float64
}

func (enm *Enemy) Update(p *Player) {
	// Calculate destination point
	destPointX := p.X
	destPointY := p.Y

	// Calculate the max distance an enemy could pass
	mx := (float64(destPointX) - float64(enm.img.Bounds().Dx())/2)
	my := (float64(destPointY) - float64(enm.img.Bounds().Dy())/2)

	// Calculate the direction vector towards the center
	dx := mx - enm.x
	dy := my - enm.y

	// Calculate the distance to the center
	distance := math.Sqrt(dx*dx + dy*dy)

	// If the image is very close to the center, stop moving
	if distance < 8 {
		gameOver = true
		return
	}

	// Normalize the direction vector
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
