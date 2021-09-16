// Package letter provides Parallel Letter Frequency solution.
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text concurrently.
func ConcurrentFrequency(words []string) FreqMap {
	result := make(FreqMap)

	ch := make(chan FreqMap, len(words))
	for _, word := range words {
		go func(s string) {
			ch <- Frequency(s)
		}(word)
	}

	for i := 0; i < len(words); i++ {
		for c, v := range <-ch {
			result[c] += v
		}
	}
	return result
}
