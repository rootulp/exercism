package foodchain

func Verse(v int) string {
	if v == 1 {
		return "I know an old lady who swallowed a fly.\nI don't know why she swallowed the fly. Perhaps she'll die."
	}
	return "I know an old lady who swallowed a spider.\nIt wriggled and jiggled and tickled inside her.\nShe swallowed the spider to catch the fly.\nI don't know why she swallowed the fly. Perhaps she'll die."
}

func Verses(start, end int) string {
	panic("Please implement the Verses function")
}

func Song() string {
	panic("Please implement the Song function")
}
