// Package protein provides Protein Translation solution.
package protein

import (
	"errors"
)

var (
	// ErrStop represents an error when a stop codon is found.
	ErrStop = errors.New("stop codon found")

	// ErrInvalidBase represents an error when an invalid base is found.
	ErrInvalidBase = errors.New("invalid codon base")
)

// FromRNA converts RNA to proteins.
func FromRNA(rna string) ([]string, error) {
	var proteins []string

	for i := 0; i < len(rna)-2; i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if errors.Is(err, ErrStop) {
			break
		}

		if err != nil {
			return proteins, err
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}

var proteinByCodon = map[string]string{
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
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon converts Codon to protein.
func FromCodon(codon string) (string, error) {
	p := proteinByCodon[codon]
	if p == "STOP" {
		return "", ErrStop
	}
	if p == "" {
		return "", ErrInvalidBase
	}
	return p, nil
}
