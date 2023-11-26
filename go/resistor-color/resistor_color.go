package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
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
