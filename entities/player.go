package entities

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	LoadTime       time.Time
	BoltAmount     int
	BoltShotBefore bool
	X              float64
	Y              float64
	Rotation       float64
	Img            *ebiten.Image // New field to store the loaded image
}

// In the Draw method of the Player struct
func (p *Player) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Get the dimensions of the image using Bounds()
	w, h := p.Img.Bounds().Dx(), p.Img.Bounds().Dy()

	// Translate to the center of the image before rotating
	opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Mirror image
	currentAngle := math.Abs(math.Mod(p.Rotation, 2*math.Pi))
	if currentAngle < math.Pi*0.5 || currentAngle > math.Pi*1.5 {
		opts.GeoM.Scale(-1, 1)
	}

	opts.GeoM.Translate(p.X, p.Y)

	// Draw the Player to the screen with the rotation options
	screen.DrawImage(p.Img, opts)
}

// Shoot method of the Player struct
func (p *Player) Shoot(g *Game) {
	// Calculate the velocity based on the Player's rotation
	vx := math.Cos(p.Rotation) * ProjectileSpeed
	vy := math.Sin(p.Rotation) * ProjectileSpeed

	// Create a new projectile with the calculated velocity
	proj := &Projectile{
		x:         p.X,
		y:         p.Y,
		rotation:  p.Rotation,
		velocityX: vx,
		velocityY: vy,
		active:    true,
		img:       g.ProjectileImg,
	}

	g.Projectiles = append(g.Projectiles, proj)
}

// Update method of the Player struct
func (p *Player) Update(g *Game) {
	// We assume one player, so id=0. For more it would be 1, then 2 and so on
	for id := range g.gamepadIDs {
		LSH := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal)
		LSV := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical)
		RSH := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickHorizontal)
		RSV := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickVertical)

		// RIght stick is aiming //to modify
		p.Rotation = math.Atan2(RSV, RSH)

		// Left stick is movement
		if math.Abs(LSH) > 0.1 {
			p.X += LSH * PlayerSpeed
		}
		if math.Abs(LSV) > 0.1 {
			p.Y += LSV * PlayerSpeed
		}

		// Handle Player shooting (enable one shot at a time and only if player has bolts available)
		FBR := ebiten.StandardGamepadButtonValue(id, ebiten.StandardGamepadButtonFrontBottomRight)
		BoltToBeShotNow := math.Abs(FBR) > 0.1
		if BoltToBeShotNow && !p.BoltShotBefore && p.BoltAmount > 0 {
			p.BoltShotBefore = true
			g.Player.Shoot(g)
			p.removeBolt()
		}
		p.BoltShotBefore = BoltToBeShotNow
	}
}

func (p *Player) addBolt() {
	if p.BoltAmount < InitialBoltAmount {
		p.BoltAmount += 1
	}
}

func (p *Player) removeBolt() {
	p.BoltAmount -= 1
}
