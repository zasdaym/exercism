package house

import (
	"fmt"
	"strings"
)

type Element struct {
	Name   string
	Action string
}

var elements = []Element{
	{},
	{},
	{"malt", "ate"},
	{"rat", "killed"},
	{"cat", "worried"},
	{"dog", "tossed"},
	{"cow with the crumpled horn", "milked"},
	{"maiden all forlorn", "kissed"},
	{"man all tattered and torn", "married"},
	{"priest all shaven and shorn", "woke"},
	{"rooster that crowed in the morn", "kept"},
	{"farmer sowing his corn", "belonged to"},
	{"horse and the hound and the horn", ""},
}

func Verse(v int) string {
	if v == 1 {
		return "This is the house that Jack built."
	}
	var lines []string
	lines = append(lines, fmt.Sprintf("This is the %s", elements[v].Name))
	for i := v - 1; i > 1; i-- {
		lines = append(lines, fmt.Sprintf("that %s the %s", elements[i].Action, elements[i].Name))
	}
	lines = append(lines, "that lay in the house that Jack built.")
	return strings.Join(lines, "\n")
}

func Song() string {
	var verses []string
	for i := 1; i < len(elements); i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}
