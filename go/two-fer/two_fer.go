// Package twofer provides Two Fer solution.
package twofer

import "fmt"

// ShareWith returns a message of sharing.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
