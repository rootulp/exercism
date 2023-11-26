package resistorcolortrio

import (
	"fmt"
	"math"
)

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

const (
	kilo = 1_000
	mega = 1_000_000
	giga = 1_000_000_000
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	colors = trimExtraColors(colors)
	decoded := decodeColors(colors)
	ohms := getValue(decoded)
	return format(ohms)
}

func format(ohms int) string {
	if ohms > giga {
		return fmt.Sprintf("%v gigaohms", ohms/giga)
	}
	if ohms > mega {
		return fmt.Sprintf("%v megaohms", ohms/mega)
	}
	if ohms > kilo {
		return fmt.Sprintf("%v kiloohms", ohms/kilo)
	}
	return fmt.Sprintf("%v ohms", ohms)
}

func trimExtraColors(colors []string) (trimmed []string) {
	// only the first 3 colors are supported
	return colors[:3]
}

func getValue(resistances []int) (value int) {
	prefix := resistances[0]*10 + resistances[1]
	numberOfZeros := resistances[2]
	return prefix * int(math.Pow10(numberOfZeros))
}

func decodeColors(colors []string) (decoded []int) {
	for _, color := range colors {
		d := decode(color)
		decoded = append(decoded, d)
	}
	return decoded
}

func decode(color string) (decoded int) {
	return colorToResistance[color]
}
