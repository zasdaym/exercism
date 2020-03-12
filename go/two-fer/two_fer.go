// Solution for Two Fer problem.
package twofer

// ShareWith will returns 'One for [name], one for me.' given a name, it fhe name is blank, name will be default to 'you'
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + subject + ", one for me."
}
