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
	if len(decoded) != 3 {
		panic("decoded should have len 3")
	}
	value := getValue(decoded)
	return fmt.Sprintf("%v ohms", value)
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
