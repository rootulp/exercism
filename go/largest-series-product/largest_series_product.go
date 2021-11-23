package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (largestProduct int64, e error) {
	if span < 0 {
		return -1, errors.New("span must not be negative")
	}
	if span > len(digits) {
		return -1, errors.New("span must be larger than digits length")
	}
	for i := 0; i <= len(digits)-span; i++ {
		product, err := getProduct(digits[i : i+span])
		if err != nil {
			return -1, err
		}
		if product > largestProduct {
			largestProduct = product
		}
	}
	return largestProduct, nil
}

func getProduct(digits string) (product int64, e error) {
	product = 1
	for _, digit := range digits {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			return -1, fmt.Errorf("cannot convert %v to int", digit)
		}
		product *= int64(d)
	}
	return product, nil
}
