package entities

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	x, y                 float64
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
	fmt.Println("Update projectile", p.x, p.y)
	// Move the projectile based on its velocity
	p.x += p.velocityX
	p.y += p.velocityY
}

func checkCollision(p *Projectile, enm *Enemy) bool {
	// Repeated use below
	px, py := p.x, p.y

	// Enemy width and height (modify if resizing the image)
	ow, oh := float64(enm.img.Bounds().Dx()), float64(enm.img.Bounds().Dy())

	return px > enm.x && px < enm.x+ow && py > enm.y && py < enm.y+oh
}

func checkPickup(prj *Projectile, plr *Player) bool {
	// Repeated use below
	prjx, prjy := prj.x, prj.y
	plrx, plry := plr.X, plr.Y

	// Player width and height (modify if resizing the image)
	ow, oh := float64(plr.Img.Bounds().Dx()), float64(plr.Img.Bounds().Dy())

	return prjx > plrx && prjx < plrx+ow && prjy > plry && prjy < plry+oh
}
