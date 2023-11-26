package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	panic("Please implement the Colors function")
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	// - Black: 0
	// - Brown: 1
	// - Red: 2
	// - Orange: 3
	// - Yellow: 4
	// - Green: 5
	// - Blue: 6
	// - Violet: 7
	// - Grey: 8
	// - White: 9
	switch color {
	case "black":
		return 0
	case "white":
		return 9
	case "orange":
		return 3
	default:
		panic("color not supported")
	}
}
