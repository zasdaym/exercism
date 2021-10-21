// Package strand provides RNA Transcription solution.
package strand

import "strings"

// ToRNA converts given DNA into RNA.
func ToRNA(dna string) string {
	var sb strings.Builder
	for _, c := range dna {
		switch c {
		case 'A':
			sb.WriteByte('U')
		case 'C':
			sb.WriteByte('G')
		case 'G':
			sb.WriteByte('C')
		case 'T':
			sb.WriteByte('A')
		}
	}
	return sb.String()
}
