// Package protein is solution for problem Protein Translation
package protein

import "errors"

var translations = map[string]string{
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

// ErrStop will throw when try to translate a stop codon
var ErrStop = errors.New("stop codon")

// ErrInvalidBase will throw when try to translate an invalid codon
var ErrInvalidBase = errors.New("invalid codon")

// FromCodon will translate given codon string to protein string
func FromCodon(s string) (string, error) {
	if _, ok := translations[s]; !ok {
		return "", ErrInvalidBase
	}

	codon := translations[s]
	if codon == "STOP" {
		return "", ErrStop
	}

	return codon, nil
}

// FromRNA will translate given RNA string to slice of protein string
func FromRNA(s string) ([]string, error) {
	codons := chunkString(s, 3)
	var result []string
	for _, codon := range codons {
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return result, err
		}

		if err == ErrStop {
			break
		}

		result = append(result, protein)
	}

	return result, nil
}

func chunkString(s string, n int) []string {
	if n >= len(s) {
		return []string{s}
	}

	var chunks []string
	chunk := make([]rune, n)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == n {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}

	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}

	return chunks
}
