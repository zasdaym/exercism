// Package letter is solution for problem Parallel Letter Frequency.
package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type freqMapMutex struct {
	freqMap FreqMap
	sync.Mutex
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the Frequency of each rune in a given text concurrently.
func ConcurrentFrequency(texts []string) FreqMap {
	m := freqMapMutex{
		freqMap: FreqMap{},
	}
	var wg sync.WaitGroup
	for _, text := range texts {
		wg.Add(1)
		go countLetter(text, &m, &wg)
	}
	wg.Wait()
	return m.freqMap
}

func countLetter(text string, m *freqMapMutex, wg *sync.WaitGroup) {
	for _, r := range text {
		m.Lock()
		m.freqMap[r]++
		m.Unlock()
	}
	wg.Done()
}
