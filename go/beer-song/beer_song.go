package beer

import "fmt"

func Song() string {
	result, err := Verses(99, 0)
	if err != nil {
		panic(err)
	}
	return result
}

func Verses(start, stop int) (output string, err error) {
	if stop < 0 {
		return "", fmt.Errorf("invalid: stop verse number %d can not be negative", stop)
	}
	if start < stop {
		return "", fmt.Errorf("invalid: start %v less than stop %v", start, stop)
	}
	for i := start; i >= stop; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		output += verse
		output += "\n"
	}
	return output, nil
}

func Verse(n int) (string, error) {
	if n > 99 {
		return "", fmt.Errorf("invalid verse number: %d", n)
	}
	switch n {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}
