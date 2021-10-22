package dna

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
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
