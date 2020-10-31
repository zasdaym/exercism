// Package twelve is solution for problem Twelve Days.
package twelve

import (
	"fmt"
	"strings"
)

type Lyric struct {
	day  string
	text string
}

var lyrics = []Lyric{
	{
		day:  "first",
		text: "a Partridge in a Pear Tree",
	},
	{
		day:  "second",
		text: "two Turtle Doves",
	},
	{
		day:  "third",
		text: "three French Hens",
	},
	{
		day:  "fourth",
		text: "four Calling Birds",
	},
	{
		day:  "fifth",
		text: "five Gold Rings",
	},
	{
		day:  "sixth",
		text: "six Geese-a-Laying",
	},
	{
		day:  "seventh",
		text: "seven Swans-a-Swimming",
	},
	{
		day:  "eighth",
		text: "eight Maids-a-Milking",
	},
	{
		day:  "ninth",
		text: "nine Ladies Dancing",
	},
	{
		day:  "tenth",
		text: "ten Lords-a-Leaping",
	},
	{
		day:  "eleventh",
		text: "eleven Pipers Piping",
	},
	{
		day:  "twelfth",
		text: "twelve Drummers Drumming",
	},
}

// Song returns a full twelve days song.
func Song() string {
	result := make([]string, 0)
	for i := range lyrics {
		result = append(result, Verse(i+1))
	}
	return strings.Join(result, "\n")
}

// Verse returns from 1 to n-th verses of the twelve days song.
func Verse(n int) string {
	var result strings.Builder
	begin := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", lyrics[n-1].day)
	result.WriteString(begin)
	if n == 1 {
		s := fmt.Sprintf("%s.", lyrics[0].text)
		result.WriteString(s)
		return result.String()
	}
	for i := n - 1; i >= 0; i-- {
		s := fmt.Sprintf("%s, ", lyrics[i].text)
		if i == 0 {
			s = fmt.Sprintf("and %s.", lyrics[i].text)
		}
		result.WriteString(s)
	}
	return result.String()
}
