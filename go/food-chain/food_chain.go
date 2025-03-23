package foodchain

import (
	"fmt"
	"strings"
)

func Verse(v int) string {
	if v == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}

	var builder strings.Builder
	builder.WriteString(openingVerse(v))
	builder.WriteString("\n")

	for i := v; i > 1; i-- {
		builder.WriteString(swallowVerse(i))
		builder.WriteString("\n")
	}

	builder.WriteString(ending)
	return builder.String()
}

func Verses(start, end int) string {
	var parts []string
	for i := start; i <= end; i++ {
		parts = append(parts, Verse(i))
	}
	return strings.Join(parts, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}

type Animal struct {
	Name   string
	Effect string
}

var animals = []Animal{
	{},
	{Name: "fly", Effect: ""},
	{Name: "spider", Effect: "It wriggled and jiggled and tickled inside her."},
	{Name: "bird", Effect: "How absurd to swallow a bird!"},
	{Name: "cat", Effect: "Imagine that, to swallow a cat!"},
	{Name: "dog", Effect: "What a hog, to swallow a dog!"},
	{Name: "goat", Effect: "Just opened her throat and swallowed a goat!"},
	{Name: "cow", Effect: "I don't know how she swallowed a cow!"},
}

func openingVerse(i int) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("I know an old lady who swallowed a %s.", animals[i].Name))
	if i > 1 {
		builder.WriteString("\n")
		builder.WriteString(animals[i].Effect)
	}
	return builder.String()
}

func swallowVerse(i int) string {
	result := fmt.Sprintf("She swallowed the %s to catch the %s", animals[i].Name, animals[i-1].Name)
	if i == 3 {
		return result + " that wriggled and jiggled and tickled inside her."
	} else {
		return result + "."
	}
}

var ending = "I don't know why she swallowed the fly. Perhaps she'll die."
