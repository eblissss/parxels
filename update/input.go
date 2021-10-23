package update

import (
	"strconv"
	"syscall/js"
)

// Globals
var g_x, g_y int
var g_kind byte = 1
var g_mousedown = false

// Add a pixel to the grid
func AddPixel(grid *[][]byte, x int, y int, kind byte) {
	gridLen := len(*grid)
	if x < gridLen && y < gridLen && x >= 0 && y >= 0 {
		(*grid)[x][gridLen-1-y] = kind
	}
}

// Update x and y position of the mouse
func getMouse(this js.Value, args []js.Value) interface{} {
	rawX := args[0].Get("clientX")
	offX := this.Get("offsetLeft")
	rawY := args[0].Get("clientY")
	offY := this.Get("offsetTop")

	g_x = (rawX.Int() - offX.Int()) / 2
	g_y = (rawY.Int() - offY.Int()) / 2

	return nil
}

func mouseMove(this js.Value, args []js.Value) interface{} {
	getMouse(this, args)
	return nil
}

func mouseUp(this js.Value, args []js.Value) interface{} {
	g_mousedown = false
	return nil
}

func mouseDown(this js.Value, args []js.Value) interface{} {
	g_mousedown = true
	return nil
}

// Check which button clicked and change kind accordingly
func changeKind(this js.Value, args []js.Value) interface{} {
	buttonID := this.Get("id").String()
	if buttonID == "airB" {
		g_kind = 0
	} else if buttonID == "sandB" {
		g_kind = 1
	} else if buttonID == "waterB" {
		g_kind = 2
	} else if buttonID == "barrierB" {
		g_kind = 3
	} else if buttonID == "smokeB" {
		g_kind = 4
	}

	return nil
}

func CheckMouse(grid *[][]byte) {
	if g_mousedown {
		AddPixel(grid, g_x, g_y, g_kind)
	}
}

func AddEvents() {
	// Add mouse events
	canvas := js.Global().Get("canvas")
	canvas.Call("addEventListener", "mousemove", js.FuncOf(mouseMove))
	canvas.Call("addEventListener", "mousedown", js.FuncOf(mouseDown))
	canvas.Call("addEventListener", "mouseup", js.FuncOf(mouseUp))

	// Add button events
	doc := js.Global().Get("document")
	buttons := doc.Call("getElementsByClassName", "b")
	for i := 0; i < buttons.Get("length").Int(); i++ {
		buttons.Get(strconv.Itoa(i)).Set("onclick", js.FuncOf(changeKind))
	}
}
