package resistorcolorduo

var colorToResistance = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	colors = trimExtraColors(colors)
	resistances := decodeColors(colors)
	return combine(resistances)
}

func trimExtraColors(colors []string) []string {
	// ignore extra colors because only two are supported
	return colors[:2]
}

func combine(resistances []int) int {
	if resistances[0] == 0 {
		return resistances[1]
	}
	return resistances[0]*10 + resistances[1]
}

func decodeColors(colors []string) (resistances []int) {
	for _, color := range colors {
		resistance := decode(color)
		resistances = append(resistances, resistance)
	}
	return resistances
}

func decode(color string) (resistance int) {
	return colorToResistance[color]
}
