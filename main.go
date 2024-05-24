package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	characterSpriteSheetPath = "sprites/rogues.png"
	itemSpriteSheetPath      = "sprites/items.png"
	playerImagePath          = "face.png"
	projectileSpeed          = 10
	rotationSpeed            = 0.1
	maxSpeedX                = 0.3
)

type Game struct {
	img                  *ebiten.Image
	objects              []*Object
	projectiles          []*Projectile
	player               *Player
	centerX              float64
	centerY              float64
	maxSpeed             float64
	minSpeed             float64
	spawnTime            time.Time
	projectilesDestroyed int
}

type Object struct {
	img     *ebiten.Image
	x, y    float64
	centerX float64
	centerY float64
	speed   float64
}

type Player struct {
	x        float64
	y        float64
	rotation float64
	image    *ebiten.Image // New field to store the loaded image
}

type Projectile struct {
	x, y      float64
	velocityX float64
	velocityY float64
	img       *ebiten.Image
}

func (g *Game) Update() error {
	// Update player
	g.player.Update()

	// Update all objects
	for _, obj := range g.objects {
		obj.Update()
	}

	// Update all projectiles
	for _, p := range g.projectiles {
		p.Update()
	}

	// Check for collisions between projectiles and objects
	g.checkCollisions()

	// Check if it's time to spawn a new object
	if time.Since(g.spawnTime).Seconds() >= 1 {
		g.spawnTime = time.Now()
		g.spawnNewObject()
	}

	// Handle player shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.player.Shoot(g)
	}

	return nil
}

func (obj *Object) Update() {
	// Calculate the direction vector towards the center
	dx := obj.centerX - obj.x
	dy := obj.centerY - obj.y

	// Calculate the distance to the center
	distance := math.Sqrt(dx*dx + dy*dy)

	// If the image is very close to the center, stop moving
	if distance < 1 {
		return
	}

	// Normalize the direction vector
	dirX := dx / distance
	dirY := dy / distance

	// Calculate the speed based on the distance
	speed := obj.speed + (obj.speed-obj.speed)*(distance/float64(math.Max(640, 480)))

	// Move the image towards the center
	obj.x += dirX * speed
	obj.y += dirY * speed
}

func (p *Projectile) Update() {
	// Move the projectile based on its velocity
	p.x += p.velocityX
	p.y += p.velocityY
}

func (g *Game) checkCollisions() {
	// Iterate over all projectiles and objects to check for collisions
	for i := len(g.projectiles) - 1; i >= 0; i-- {
		p := g.projectiles[i]
		for j := len(g.objects) - 1; j >= 0; j-- {
			obj := g.objects[j]
			if checkCollision(p, obj) {
				// Remove the projectile and the object
				g.projectiles = append(g.projectiles[:i], g.projectiles[i+1:]...)
				g.objects = append(g.objects[:j], g.objects[j+1:]...)
				g.projectilesDestroyed++
				break
			}
		}
	}
}

func checkCollision(p *Projectile, obj *Object) bool {
	px, py := p.x, p.y
	ox, oy := obj.x, obj.y
	ow, oh := obj.img.Size()
	ow *= 2
	oh *= 2
	return px > ox && px < ox+float64(ow) && py > oy && py < oy+float64(oh)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the counter
	// // Render the counter using the default font

	// Fill the screen with a white background
	screen.Fill(color.Black)

	// Draw all objects
	for _, obj := range g.objects {
		obj.Draw(screen)
	}

	// Draw all projectiles
	for _, p := range g.projectiles {
		p.Draw(screen)
	}

	// Draw the player
	g.player.Draw(screen)
	stringToDisplay := "Enemies destroyed: " + strconv.Itoa(g.projectilesDestroyed)
	ebitenutil.DebugPrint(screen, stringToDisplay)
}

func (obj *Object) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Scale the image to be twice its original size
	opts.GeoM.Scale(2, 2)

	// Set the position of the image
	opts.GeoM.Translate(obj.x, obj.y)

	// Draw the image to the screen with the scaling options
	screen.DrawImage(obj.img, opts)
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Set the position of the projectile
	opts.GeoM.Translate(p.x, p.y)

	// Draw the projectile to the screen
	screen.DrawImage(p.img, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Define the game's screen size.
	return 640, 480
}

func (g *Game) spawnNewObject() {
	// Create a new Object
	obj := &Object{
		img:     g.img,
		speed:   g.maxSpeed,
		centerX: g.centerX,
		centerY: g.centerY,
	}

	// Randomly choose an edge (0=left, 1=top, 2=right, 3=bottom)
	edge := rand.Intn(4)
	switch edge {
	case 0:
		// Left edge
		obj.x = -float64(obj.img.Bounds().Dx()) * 2
		obj.y = rand.Float64() * 480
	case 1:
		// Top edge
		obj.x = rand.Float64() * 640
		obj.y = -float64(obj.img.Bounds().Dy()) * 2
	case 2:
		// Right edge
		obj.x = 640
		obj.y = rand.Float64() * 480
	case 3:
		// Bottom edge
		obj.x = rand.Float64() * 640
		obj.y = 480
	}

	// Add the new object to the objects slice
	g.objects = append(g.objects, obj)
}

// In the Draw method of the Player struct
func (p *Player) Draw(screen *ebiten.Image) {
	// Create a new DrawImageOptions struct
	opts := &ebiten.DrawImageOptions{}

	// Get the dimensions of the image using Bounds()
	w, h := p.image.Bounds().Dx(), p.image.Bounds().Dy()

	// Translate to the center of the image before rotating
	opts.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate the player around its center
	opts.GeoM.Rotate(p.rotation)

	// Set the position of the player
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
		velocityX: vx,
		velocityY: vy,
		img:       g.img,
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

func loadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func loadCharacterFromSheet(sheet *ebiten.Image, index int) *ebiten.Image {
	spriteSize := 32

	// Calculate the x coordinate of the character in the sprite sheet
	x := index * spriteSize

	// Define the rectangle for the character (assuming they are in one "line")
	rect := image.Rect(x, 0, x+spriteSize, spriteSize)

	// Use SubImage to get the specific character
	subImage := sheet.SubImage(rect).(*ebiten.Image)

	return subImage
}

func main() {

	game := &Game{
		maxSpeed:  maxSpeedX, // Maximum speed at the edges
		minSpeed:  0,         // Minimum speed at the center
		spawnTime: time.Now(),
	}

	var err error
	game.img, err = loadImage("X1.png") // Use the item sheet
	if err != nil {
		log.Fatal(err)
	}

	// Load the player image from the sprite sheet
	characterSheet, _, err := ebitenutil.NewImageFromFile(characterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Extract the character at index 0 (first character)
	character := loadCharacterFromSheet(characterSheet, 0)

	// Calculate the center of the screen
	game.centerX = (640 - float64(game.img.Bounds().Dx())*2) / 2
	game.centerY = (480 - float64(game.img.Bounds().Dy())*2) / 2

	// Initialize the player with the loaded image
	game.player = &Player{
		x:        320,
		y:        240,
		rotation: 0,         // Initial rotation angle
		image:    character, // Assign the loaded player image
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Shoot the ghosts!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
