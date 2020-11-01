// Package twelve is solution for problem Twelve Days.
package twelve

import (
	"fmt"
	"strings"
)

// Lyric represents a lyric part in twelve days song.
type Lyric struct {
	day  string
	text string
}

var lyrics = []Lyric{
	{
		day:  "first",
		text: "a Partridge in a Pear Tree.",
	},
	{
		day:  "second",
		text: "two Turtle Doves, and ",
	},
	{
		day:  "third",
		text: "three French Hens, ",
	},
	{
		day:  "fourth",
		text: "four Calling Birds, ",
	},
	{
		day:  "fifth",
		text: "five Gold Rings, ",
	},
	{
		day:  "sixth",
		text: "six Geese-a-Laying, ",
	},
	{
		day:  "seventh",
		text: "seven Swans-a-Swimming, ",
	},
	{
		day:  "eighth",
		text: "eight Maids-a-Milking, ",
	},
	{
		day:  "ninth",
		text: "nine Ladies Dancing, ",
	},
	{
		day:  "tenth",
		text: "ten Lords-a-Leaping, ",
	},
	{
		day:  "eleventh",
		text: "eleven Pipers Piping, ",
	},
	{
		day:  "twelfth",
		text: "twelve Drummers Drumming, ",
	},
}

// Song returns a full twelve days song.
func Song() string {
	song := make([]string, 0)
	for i := range lyrics {
		song = append(song, Verse(i+1))
	}
	return strings.Join(song, "\n")
}

// Verse returns from 1 to n-th verses of the twelve days song.
func Verse(n int) string {
	var result strings.Builder
	begin := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", lyrics[n-1].day)
	result.WriteString(begin)
	for i := n - 1; i >= 0; i-- {
		s := fmt.Sprintf("%s", lyrics[i].text)
		result.WriteString(s)
	}
	return result.String()
}
