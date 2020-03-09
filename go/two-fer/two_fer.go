// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	subject := "you"

	if name != "" {
		subject = name
	}

	return "One for " + subject + ", one for me."
}
