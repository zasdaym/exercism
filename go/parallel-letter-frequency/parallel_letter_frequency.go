// Package letter is solution for problem Parallel Letter Frequency.
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

// ConcurrentFrequency counts the Frequency of each rune in given texts concurrently.
func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap, 10)
	for _, text := range texts {
		go func(s string) {
			ch <- Frequency(s)
		}(text)
	}

	m := FreqMap{}
	for range texts {
		for k, v := range <-ch {
			m[k] += v
		}
	}
	return m
}

