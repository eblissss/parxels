package update

// Enum for kinds
const (
	Air byte = iota
	Sand
	Water
	Barrier
	Smoke

	// Mark to lock for current update
	Mark byte = 128
)

// Update all kinds
func UpdateAll(grid *[][]byte) {
	gridLen := len(*grid)
	for x := 0; x < gridLen; x++ {
		for y := 0; y < gridLen; y++ {
			if (*grid)[x][y] == Sand {
				updateSand(grid, x, y)
			} else if (*grid)[x][y] == Water {
				updateWater(grid, x, y)
			} else if (*grid)[x][y] == Smoke {
				updateSmoke(grid, x, y)
			}
		}
	}
	// Undo marks
	for x := 0; x < gridLen; x++ {
		for y := 0; y < gridLen; y++ {
			if (*grid)[x][y] >= Mark {
				(*grid)[x][y] -= Mark
			}
		}
	}
}

// Swap itself with something lower
func gravity(grid *[][]byte, x int, y int, kind byte, swap byte) bool {
	if y > 0 && (*grid)[x][y-1] == swap {
		if swap != Air {
			swap += Mark
		}
		(*grid)[x][y-1] = kind + Mark
		(*grid)[x][y] = swap

		return true
	}
	return false
}

// Swap itself with air (reversed)
func revGravity(grid *[][]byte, x int, y int, kind byte) bool {
	if y < len(*grid)-1 && (*grid)[x][y+1] == Air {
		(*grid)[x][y+1] = kind + Mark
		(*grid)[x][y] = Air
		return true
	}
	return false
}
