// Solution for problem Scale Generator.
package scale

import "strings"

// Scale will return musical scale given starting note and intervals.
func Scale(tonic, interval string) (scale []string) {
	sharpScales := []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

	flatScales := []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		scale = sharpScales
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = flatScales
	}

	tonic = strings.Title(tonic)

	for i, elem := range scale {
		if elem == tonic {
			scale = append(scale[i:], scale[:i]...)
			break
		}
	}

	if interval == "" {
		return scale
	}

	partialScale := []string{}

	stepSize := map[string]int{
		"m": 1,
		"M": 2,
		"A": 3,
	}

	i := 0

	for _, diff := range strings.Split(interval, "") {
		if step, ok := stepSize[diff]; ok {
			partialScale = append(partialScale, scale[i%len(scale)])
			i += step
		}
	}

	return partialScale
}
