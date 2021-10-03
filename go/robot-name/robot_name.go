// Package robotname provides Robot Name solution.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	limit     = 26 * 26 * 10 * 10 * 10
	maxNumber = 1000
	maxLetter = 26
)

// namePool contains all possible robot name.
var namePool = generate()

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// generate generates all possible robot name.
func generate() []string {
	var names []string

	for i := 0; i < limit; i++ {
		num := i % maxNumber
		letterPart := i / maxNumber
		first := letterPart/maxLetter + 'A'
		second := letterPart%maxLetter + 'A'
		name := fmt.Sprintf("%c%c%03d", first, second, num)
		names = append(names, name)
	}

	random.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	return names
}

// used tracks how many robot name has been used.
var used int

// Robot represents a robot with a randomly generated unique name.
type Robot struct {
	name string
}

// Name returns the robot's name and generates a new one if it's not already set.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if used == limit {
		return "", errors.New("no more name can be generated")
	}

	r.name = namePool[used]
	used++

	return r.name, nil
}

// Reset clears the robot's name.
func (r *Robot) Reset() {
	r.name = ""
}
