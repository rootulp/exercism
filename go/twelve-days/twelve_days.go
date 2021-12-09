package twelve

import (
	"fmt"
	"strings"
)

func Verse(day int) string {
	return fmt.Sprintf("On the %v day of Christmas my true love gave to me: %s.", ordinal(day), clause(day))
}

func ordinal(day int) string {
	dayToOrdinal := map[int]string{
		1:  "first",
		2:  "second",
		3:  "third",
		4:  "fourth",
		5:  "fifth",
		6:  "sixth",
		7:  "seventh",
		8:  "eighth",
		9:  "ninth",
		10: "tenth",
		11: "eleventh",
		12: "twelfth",
	}
	return dayToOrdinal[day]
}

func clause(day int) (result string) {
	clauses := []string{}
	for i := day; i >= 1; i-- {
		if day > 1 && i == 1 {
			clauses = append(clauses, fmt.Sprintf("and %s", gift(i)))
		} else {
			clauses = append(clauses, gift(i))
		}
	}
	result = strings.Join(clauses, ", ")
	return result
}

func gift(day int) string {
	dayToGift := map[int]string{
		12: "twelve Drummers Drumming",
		11: "eleven Pipers Piping",
		10: "ten Lords-a-Leaping",
		9:  "nine Ladies Dancing",
		8:  "eight Maids-a-Milking",
		7:  "seven Swans-a-Swimming",
		6:  "six Geese-a-Laying",
		5:  "five Gold Rings",
		4:  "four Calling Birds",
		3:  "three French Hens",
		2:  "two Turtle Doves",
		1:  "a Partridge in a Pear Tree",
	}
	return dayToGift[day]
}

func Song() string {
	panic("Please implement the Song function")
}
