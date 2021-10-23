package update

func updateWater(grid *[][]byte, x int, y int) {
	gridLen := len(*grid)

	// Air interactions
	if gravity(grid, x, y, Water, Air) {
	} else {
		findEmpty(grid, gridLen, x, y)
	}
}

func findEmpty(grid *[][]byte, gridLen int, x int, y int) bool {
	if y == 0 {
		return false
	}
	goRight, goLeft := true, true

	for i := 1; i < gridLen; i++ {
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
				(*grid)[leftI][y-1] = Water + Mark
				(*grid)[x][y] = Air
				return true
			}
		}
		if goRight {
			if (*grid)[rightI][y-1] == Air {
				(*grid)[rightI][y-1] = Water + Mark
				(*grid)[x][y] = Air
				return true
			}
		}
	}
	return false
}
