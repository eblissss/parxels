package main

// Parxels - Particle Pixels
// Ethan Bliss
// October 2021

// A WebAssembly app with an interactive grid where the user can
// place different kinds of particle pixels that interact with
// each other and follow simple rules.

// To fix: large wasm file size
// Tinygo breaks the program :(

import (
	"syscall/js"
	"time"

	"github.com/eblissss/parxels/display"
	"github.com/eblissss/parxels/update"
)

// Make a size x size grid (0, 0) TL
func createGrid(size int) [][]byte {
	grid := make([][]byte, size)
	for i := range grid {
		grid[i] = make([]byte, size)
	}
	return grid
}

func main() {
	// Get canvas context and id
	canvas := js.Global().Get("canvas")
	ctx := canvas.Call("getContext", "2d")
	cWidth := canvas.Get("width").Int()
	cHeight := canvas.Get("height").Int()
	id := ctx.Call("getImageData", 0, 0, cWidth, cHeight)

	grid := createGrid(cWidth)

	/*for i := 0; i < 80; i++ {
		for j := 0; j < 80; j += 2 {
			update.AddPixel(&grid, i, j, 1)
		}
	}*/

	// Add DOM events
	update.AddEvents()

	// Main loop
	for true {
		time.Sleep(time.Millisecond * 5)
		update.CheckMouse(&grid)

		time.Sleep(time.Millisecond * 5)
		update.CheckMouse(&grid)

		update.UpdateAll(&grid)
		display.DisplayOnCanvas(grid, id, ctx)
	}
}
