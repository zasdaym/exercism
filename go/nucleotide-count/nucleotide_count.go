// Package dna provides Nucleotide Count solution.
package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
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

	for _, c := range d {
		if _, ok := h[c]; !ok {
			return h, fmt.Errorf("invalid nucleotide: %c", c)
		}

		h[c]++
	}

	return h, nil
}
