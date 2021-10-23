package update

func updateSmoke(grid *[][]byte, x int, y int) {
	gridLen := len(*grid)

	// Upwards motion (air)
	if revGravity(grid, x, y, Smoke) {
		// Upwards motion general
	} else if y < gridLen-1 && (*grid)[x][y+1] != Barrier && (*grid)[x][y+1] != Smoke {
		tempKind := (*grid)[x][y+1]
		(*grid)[x][y+1] = Smoke + Mark
		(*grid)[x][y] = tempKind + Mark
	} else {
		// Spread out upwards
		findEmptyTop(grid, gridLen, x, y)
	}
}

// Find the closest empty spot in the row above
func findEmptyTop(grid *[][]byte, gridLen int, x int, y int) bool {
	if y == gridLen-1 {
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

		// Bounds check
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

		// If space found, take it
		if goLeft {
			if (*grid)[leftI][y+1] == Air {
				(*grid)[leftI][y+1] = Smoke + Mark
				(*grid)[x][y] = Air
				return true
			}
		}
		if goRight {
			if (*grid)[rightI][y+1] == Air {
				(*grid)[rightI][y+1] = Smoke + Mark
				(*grid)[x][y] = Air
				return true
			}
		}
	}
	return false
}
