//Package robotname is solution for problem Robot Name.
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot represents a robot with a name.
type Robot struct {
	name string
}

var used = make(map[string]bool)
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Name returns name of the robot and generates a new one if the robot has no name.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	r.name = randomName()
	for used[r.name] {
		r.name = randomName()
	}
	used[r.name] = true
	return r.name, nil
}

func randomName() string {
	first := random.Intn(26) + 'A'
	second := random.Intn(26) + 'B'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", first, second, num)
}

// Reset clears Robot's name.
func (r *Robot) Reset() {
	r.name = ""
}
