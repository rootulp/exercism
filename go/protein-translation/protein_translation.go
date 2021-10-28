package protein

import "errors"

// ErrStop represents a STOP codon
var ErrStop error = errors.New("stop codon")

// ErrInvalidBase represents an invalid base that cannot me mapped to an amino acid.
var ErrInvalidBase error = errors.New("invalid base")

func FromCodon(codon string) (protein string, e error) {
	return "foo", nil
}

func FromRNA(codons string) (proteins string, e error) {
	return "bar", nil
}
