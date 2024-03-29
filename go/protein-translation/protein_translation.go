package protein

import (
	"errors"
)

// ErrStop represents a stop codon
var ErrStop error = errors.New("stop codon")

// ErrInvalidBase represents an invalid base that cannot me mapped to an amino acid.
var ErrInvalidBase error = errors.New("invalid base")

var stopCodons = map[string]bool{
	"UAA": true,
	"UAG": true,
	"UGA": true,
}

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
}

func FromCodon(codon string) (protein string, e error) {
	if _, ok := stopCodons[codon]; ok {
		return "", ErrStop
	}
	if _, ok := codonToProtein[codon]; !ok {
		return "", ErrInvalidBase
	}
	return codonToProtein[codon], nil
}

const CODON_LENGTH = 3

func FromRNA(codons string) (proteins []string, e error) {
	for i := 0; i <= len(codons)-CODON_LENGTH; i += CODON_LENGTH {
		codon := codons[i : i+CODON_LENGTH]
		codon, err := FromCodon(codon)
		if errors.Is(err, ErrStop) {
			return proteins, nil
		}
		if errors.Is(err, ErrInvalidBase) {
			return proteins, ErrInvalidBase
		}
		proteins = append(proteins, codon)
	}
	return proteins, nil
}
