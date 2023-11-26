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

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	decoded := decodeColors(colors)
	ohms := getValue(decoded)
	if ohms > 1_000_000_000 {
		gigaohms := ohms / 1_000_000_000
		return fmt.Sprintf("%v gigaohms", gigaohms)
	}
	if ohms > 1_000_000 {
		megaohms := ohms / 1_000_000
		return fmt.Sprintf("%v megaohms", megaohms)
	}
	if ohms > 1000 {
		kiloohms := ohms / 1000
		return fmt.Sprintf("%v kiloohms", kiloohms)
	}
	return fmt.Sprintf("%v ohms", ohms)
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
