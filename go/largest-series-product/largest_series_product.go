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
		return -1, errors.New("error span can not be larger than the number of digits provided")
	}
	fmt.Printf("LargestSeriesProduct(%v, %v)\n", digits, span)
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
			return -1, fmt.Errorf("cannot conver %v to int", digit)
		}
		product *= int64(d)
	}
	fmt.Printf("getProduct(%v)=%v\n", digits, product)
	return product, nil
}
