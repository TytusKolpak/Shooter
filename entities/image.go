package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	spriteSize = 32
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
