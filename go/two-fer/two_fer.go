package twofer

import "fmt"

// ShareWith returns a string in the following format: "One for you, one for me."
// if a name is provided, it is used instead of "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
