package entities

import (
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadSpriteFromSheet(sheet *ebiten.Image, posX, posY int) *ebiten.Image {
	// Calculate coordinates of the sprite in the sprite sheet
	x := posX * spriteSize
	y := posY * spriteSize

	// Define the rectangle for the sprite (assuming they are in one "line")
	rect := image.Rect(x, y, x+spriteSize, y+spriteSize)

	// Use SubImage to get the specific sprite
	subImage := sheet.SubImage(rect).(*ebiten.Image)

	return subImage
}

// Draw background by repeating an image (we prepare it here and redraw later)
func GenerateBackground(sheet *ebiten.Image) *ebiten.Image {
	// Prepare large blank image to fill later
	composedImage := ebiten.NewImage(ScreenWidth, ScreenHeight)

	// Calculate the number of times the image needs to be drawn
	horizontalTiles := (ScreenWidth + spriteSize - 1) / spriteSize
	verticalTiles := (ScreenHeight + spriteSize - 1) / spriteSize

	// Choose pattern of tiles: indexes 6-15 contain different pattern set
	tileY := rand.Intn(10) + 6

	// Draw the image repeatedly to fill the screen
	for y := 0; y < verticalTiles; y++ {
		for x := 0; x < horizontalTiles; x++ {
			opts := &ebiten.DrawImageOptions{}
			// Draw this new tile at given position
			opts.GeoM.Translate(float64(x*spriteSize), float64(y*spriteSize))
			tileX := rand.Intn(3) + 1

			// Use random one of the given pattern tiles
			composedImage.DrawImage(LoadSpriteFromSheet(sheet, tileX, tileY), opts)
		}
	}

	return composedImage
}
