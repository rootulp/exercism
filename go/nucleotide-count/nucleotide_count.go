package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	if !d.isValid() {
		return Histogram{}, fmt.Errorf("DNA stand %s contains invalid nucleotides", d)
	}
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, r := range d {
		h[r] += 1
	}
	return h, nil
}

func (d DNA) isValid() bool {
	for _, r := range d {
		if r == 'A' || r == 'C' || r == 'G' || r == 'T' {
			continue
		} else {
			return false
		}
	}
	return true
}
