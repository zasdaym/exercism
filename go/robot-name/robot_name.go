//Package robotname is solution for problem Robot Name.
package robotname

import (
	"fmt"
)

// Robot represents a robot with a name.
type Robot struct {
	name string
}

// Name returns name of the robot and generates a new one if the robot has no name.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	name, err := randomName()
	if err != nil {
		return "", err
	}
	r.name = name
	return r.name, nil
}

var counter = 0

func randomName() (string, error) {
	if counter > 675999 {
		return "", fmt.Errorf("all combinations used")
	}
	counter++
	return (generateLetters(counter) + generateNumber(counter)), nil
}

func generateLetters(counter int) string {
	n := (counter / 1000) + 65
	first := string(rune(n / 65 + 65))
	second := string(rune(n % 65 + 65))
	return first + second
}

func generateNumber(counter int) string {
	n := (counter % 1000)
	number := fmt.Sprintf("%03d", n)
	return number
}

// Reset clears Robot's name.
func (r *Robot) Reset() {
	r.name = ""
}
