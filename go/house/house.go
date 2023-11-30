package house

import (
	"fmt"
	"strings"
)

func Verse(v int) string {
	if v == 1 {
		return `This is the house that Jack built.`
	}
	if v == 2 {
		return `This is the malt
that lay in the house that Jack built.`
	}
	if v == 3 {
		return `This is the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 4 {
		return `This is the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 5 {
		return `This is the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 6 {
		return `This is the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 7 {
		return `This is the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 8 {
		return `This is the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 9 {
		return `This is the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 10 {
		return `This is the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 11 {
		return `This is the farmer sowing his corn
that kept the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	if v == 12 {
		return `This is the horse and the hound and the horn
that belonged to the farmer sowing his corn
that kept the rooster that crowed in the morn
that woke the priest all shaven and shorn
that married the man all tattered and torn
that kissed the maiden all forlorn
that milked the cow with the crumpled horn
that tossed the dog
that worried the cat
that killed the rat
that ate the malt
that lay in the house that Jack built.`
	}
	panic(fmt.Sprintf("unsupported verseNumber %v", v))
}

func Song() string {
	verses := []string{}
	for i := 1; i <= 12; i++ {
		verse := Verse(i)
		verses = append(verses, verse)
	}
	return strings.Join(verses, "\n\n")
}
