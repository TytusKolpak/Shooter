package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	maxEnemySpeed = 2
	minEnemySpeed = 0.2
)

type Enemy struct {
	img  *ebiten.Image
	x, y float64
}

func (enm *Enemy) Update() {
	// Calculate the max distance an enemy could pass
	mx := (screenWidth/2 - float64(enm.img.Bounds().Dx())/2)
	my := (screenHeight/2 - float64(enm.img.Bounds().Dy())/2)
	maxDistance := math.Sqrt(mx*mx + my*my)

	// Calculate the direction vector towards the center
	dx := mx - enm.x
	dy := my - enm.y

	// Calculate the distance to the center
	distance := math.Sqrt(dx*dx + dy*dy)

	// If the image is very close to the center, stop moving
	if distance < 1 {
		return
	}

	// Normalize the direction vector
	dirX := dx / distance
	dirY := dy / distance

	// Move the image towards the center (slow down the closer they are)
	enm.x += dirX * maxEnemySpeed * (distance/maxDistance + minEnemySpeed)
	enm.y += dirY * maxEnemySpeed * (distance/maxDistance + minEnemySpeed)
}

func (enm *Enemy) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Set the position of the image
	opts.GeoM.Translate(enm.x, enm.y)

	// Draw the image to the screen with the scaling options
	screen.DrawImage(enm.img, opts)
}
