package lsproduct

import (
	"fmt"
	"log"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (largestProduct int64, err error) {
	fmt.Printf("LargestSeriesProduct(%v, %v)\n", digits, span)
	for i := 0; i <= len(digits)-span; i++ {
		product := getProduct(digits[i : i+span])
		if product > largestProduct {
			largestProduct = product
		}
	}
	return largestProduct, nil
}

func getProduct(digits string) (product int64) {
	product = 1
	for _, digit := range digits {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			log.Fatalf("cannot convert %v to int", digit)
		}
		product *= int64(d)
	}
	fmt.Printf("getProduct(%v)=%v\n", digits, product)
	return product
}
