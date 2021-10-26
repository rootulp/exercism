package strand

var dnaToRna = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	result := ""
	for _, r := range dna {
		result += string(dnaToRna[r])
	}
	return result
}
