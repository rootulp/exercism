package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	} else if len(rhyme) == 1 {
		return []string{footer(rhyme[0])}
	}
	return multilineProverb(rhyme)
}

func multilineProverb(rhyme []string) (result []string) {
	for i, item := range rhyme {
		if i == len(rhyme)-1 {
			result = append(result, footer(rhyme[0]))
		} else {
			result = append(result, verse(item, rhyme[i+1]))
		}
	}
	return result
}

func verse(itemA string, itemB string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", itemA, itemB)

}

func footer(item string) string {
	return fmt.Sprintf("And all for the want of a %s.", item)
}
