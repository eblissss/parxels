package display

import (
	"syscall/js"
)

var r, g, b, a byte = 0, 0, 0, 0

// Get the corresponding RGB color from specified kind
func pickColor(pixel byte) (byte, byte, byte, byte) {
	if pixel == 1 {
		r, g, b, a = 200, 190, 133, 255
	} else if pixel == 2 {
		r, g, b, a = 6, 67, 200, 255
	} else if pixel == 3 {
		r, g, b, a = 2, 46, 46, 255
	} else if pixel == 4 {
		r, g, b, a = 80, 80, 80, 255
	} else {
		a = 0
	}
	return r, g, b, a
}

// Draw grid on canvas
func DisplayOnCanvas(grid [][]byte, id js.Value, ctx js.Value) {
	gridSize := len(grid)
	pixels := make([]byte, gridSize*gridSize*4)

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			// Rotate (grid grows up, canvas grows down)
			invJ := gridSize - 1 - j
			off := (invJ*gridSize + i) * 4
			pickColor(grid[i][j])
			pixels[off] = r
			pixels[off+1] = g
			pixels[off+2] = b
			pixels[off+3] = a
		}
	}

	// Push image to canvas
	idData := id.Get("data")
	js.CopyBytesToJS(idData, pixels)
	ctx.Call("putImageData", id, 0, 0)
}
