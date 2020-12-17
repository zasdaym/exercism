// Package beer is solution for problem Beer Song.
package beer

import (
	"fmt"
)

var template string = "%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n"

// Verse returns a specific verse.
func Verse(n int) (string, error) {
	if n == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	if n == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if n == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}
	if n < 0 || n > 100 {
		return "", fmt.Errorf("verse number must be between 1 and 100")
	}
	return fmt.Sprintf(template, n, n, n-1), nil
}

// Verses returns a range of verses.
func Verses(upper, lower int) (string, error) {
	if upper < lower {
		return "", fmt.Errorf("invalid start and stop verse number")
	}
	var verses string
	for i := upper; i > lower-1; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses += verse
		verses += "\n"
	}
	return verses, nil
}

// Song returns full Beer Song.
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
