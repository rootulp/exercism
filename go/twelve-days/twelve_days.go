package twelve

import "fmt"

func Verse(i int) string {
	return fmt.Sprintf("On the %v day of Christmas my true love gave to me: a Partridge in a Pear Tree.", ordinal(i))
}

func ordinal(i int) string {
	intToOrdinal := map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eigth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
	}
	return intToOrdinal[i]
}

func Song() string {
	panic("Please implement the Song function")
}
