package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadSpriteFromSheet(sheet *ebiten.Image, posX, posY int) *ebiten.Image {
	spriteSize := 32

	// Calculate coordinates of the sprite in the sprite sheet
	x := posX * spriteSize
	y := posY * spriteSize

	// Define the rectangle for the sprite (assuming they are in one "line")
	rect := image.Rect(x, y, x+spriteSize, y+spriteSize)

	// Use SubImage to get the specific sprite
	subImage := sheet.SubImage(rect).(*ebiten.Image)

	return subImage
}
