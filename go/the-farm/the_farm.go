package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if fodderAmount < 0 {
		// This exercism problem is flawed because it expects a capitalized
		// error string which violate go lint
		return 0, errors.New("Negative fodder")
	}
	if err == ErrScaleMalfunction {
		return fodderAmount * 2 / float64(cows), nil
	}
	if err != nil {
		return 0, err
	}
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, fmt.Errorf("silly nephew, there cannot be %d cows", cows)
	}
	return fodderAmount / float64(cows), nil
}
