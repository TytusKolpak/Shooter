package entities

import (
	"image"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func LoadSpriteFromSheet(sheet *ebiten.Image, posX, posY int) *ebiten.Image {
	// Calculate coordinates of the sprite in the sprite sheet
	x := posX * spriteSize
	y := posY * spriteSize

	// Define the rectangle for the sprite
	rect := image.Rect(x, y, x+spriteSize, y+spriteSize)

	// Use SubImage to get the specific sprite
	subImage := sheet.SubImage(rect).(*ebiten.Image)

	return subImage
}

func AddBoundingBox(image *ebiten.Image) *ebiten.Image {

	// Create a new image to draw the sprite and bounding box
	boundingBoxImg := ebiten.NewImage(spriteSize, spriteSize)

	// Draw the sprite onto the new image
	op := &ebiten.DrawImageOptions{}
	boundingBoxImg.DrawImage(image, op)

	// Draw the bounding box onto the new image (as 4 lines) x, y, width, height
	vector.DrawFilledRect(boundingBoxImg, 0, 0, spriteSize, 1, color.RGBA{255, 0, 0, 255}, true)            // top
	vector.DrawFilledRect(boundingBoxImg, 0, 0, 1, spriteSize, color.RGBA{255, 0, 0, 255}, true)            // left
	vector.DrawFilledRect(boundingBoxImg, 0, spriteSize-1, spriteSize, 1, color.RGBA{255, 0, 0, 255}, true) // bottom
	vector.DrawFilledRect(boundingBoxImg, spriteSize-1, 0, 1, spriteSize, color.RGBA{255, 0, 0, 255}, true) // right

	return boundingBoxImg
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

func LoadSpriteSheets() (*ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image) {
	// Load the background image (put these in entities somewhere)
	tileSheet, _, err := ebitenutil.NewImageFromFile(TileSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the item image from the sprite sheet
	enemySheet, _, err := ebitenutil.NewImageFromFile(MonsterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the item image from the sprite sheet
	itemSheet, _, err := ebitenutil.NewImageFromFile(ItemSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load the Player image from the sprite sheet
	characterSheet, _, err := ebitenutil.NewImageFromFile(CharacterSpriteSheetPath)
	if err != nil {
		log.Fatal(err)
	}

	return tileSheet, enemySheet, itemSheet, characterSheet
}
