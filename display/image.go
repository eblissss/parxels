package display

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Save the grid as an image (WIP)
func GridToImage(grid [][]byte, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, len(grid), len(grid)))
	for x, row := range grid {
		for y, coord := range row {
			imgY := len(grid) - y
			r, g, b, a := pickColor(coord)
			img.Set(x, imgY, color.RGBA{r, g, b, a})
		}
	}
	f, _ := os.Create(filename)
	defer f.Close()
	png.Encode(f, img)
}
