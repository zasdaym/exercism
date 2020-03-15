// Package strand is solution for problem RNA Transcription.
package strand

// ToRNA will return the RNA complement of given DNA strand.
func ToRNA(dna string) string {
	var result []rune
	codes := []rune("GCTA")
	complement := []rune("CGAU")

	for _, x := range dna {
		found := false
		for i := range codes {
			if x == codes[i] {
				result = append(result, complement[i])
				found = true
				break
			}
		}

		if !found {
			result = append(result, x)
		}
	}
	return string(result)
}
