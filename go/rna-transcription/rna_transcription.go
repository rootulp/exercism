package strand

var dnaToRna = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) (rna string) {
	for _, r := range dna {
		rna += string(dnaToRna[r])
	}
	return rna
}
