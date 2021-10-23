package update

func updateSand(grid *[][]byte, x int, y int) {
	gridLen := len(*grid)

	// Air interactions
	if gravity(grid, x, y, Sand, Air) {
	} else if sandSpread(grid, gridLen, x, y, Air) {
		// Water interactions
	} else if gravity(grid, x, y, Sand, Water) {
	} else if sandSpread(grid, gridLen, x, y, Water) {
	}
}

func sandSpread(grid *[][]byte, gridLen int, x int, y int, kind byte) bool {
	if y == 0 {
		return false
	}
	goRight, goLeft := true, true

	var max int
	if kind == Air {
		max = 2
	} else if kind == Water {
		max = 1
	}

	// Go up to two places over

	for i := 1; i < max+1; i++ {
		// Return if nowhere to go
		if !(goRight || goLeft) {
			return false
		}

		leftI := x - i
		rightI := x + i

		// Bounds checks
		if leftI < 0 {
			goLeft = false
		}
		if goLeft && (*grid)[leftI][y] != Air {
			goLeft = false
		}
		if rightI >= gridLen {
			goRight = false
		}
		if goRight && (*grid)[rightI][y] != Air {
			goRight = false
		}

		// Place if spot found
		if goLeft {
			if (*grid)[leftI][y-1] == Air {
				(*grid)[leftI][y-1] = Sand + Mark
				(*grid)[x][y] = Air
				return true
			}
		}
		if goRight {
			if (*grid)[rightI][y-1] == Air {
				(*grid)[rightI][y-1] = Sand + Mark
				(*grid)[x][y] = Air
				return true
			}
		}
	}
	return false
}
