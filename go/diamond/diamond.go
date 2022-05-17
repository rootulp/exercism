package diamond

import "fmt"

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("char %v is not a valid capital letter", char)
	}
	return "A", nil
}
