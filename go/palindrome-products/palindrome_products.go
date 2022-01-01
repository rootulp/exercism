package palindrome

import (
	"errors"
	"strconv"
)

// Define Product type here.
type Product struct {
	palindrome     int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (min Product, max Product, e error) {
	if fmin > fmax {
		// This error message doesn't conform to go-staticcheck but this problem expects this exact error string
		return min, max, errors.New("fmin > fmax...")
	}
	products := getProducts(fmin, fmax)
	if len(products) == 0 {
		// This error message doesn't conform to go-staticcheck but this problem expects this exact error string
		return min, max, errors.New("no palindromes...")
	}
	return getMin(products), getMax(products), nil
}

func getProducts(min int, max int) (products []Product) {
	palindromeToFactors := getPalindromeToFactors(min, max)
	for palindrome, factors := range palindromeToFactors {
		product := Product{
			palindrome:     palindrome,
			Factorizations: factors,
		}
		products = append(products, product)
	}
	return products
}

func getPalindromeToFactors(min int, max int) (palindromeToFactors map[int][][2]int) {
	palindromeToFactors = make(map[int][][2]int)
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			candidate := i * j
			factor := [2]int{i, j}
			if isPalindrome(candidate) {
				if factors, ok := palindromeToFactors[candidate]; ok {
					factors = append(factors, factor)
					palindromeToFactors[candidate] = factors
				} else {
					palindromeToFactors[candidate] = [][2]int{factor}
				}
			}
		}
	}
	return palindromeToFactors
}

func getMin(products []Product) (min Product) {
	min = products[0]
	for _, product := range products {
		if product.palindrome < min.palindrome {
			min = product
		}
	}
	return min
}

func getMax(products []Product) (max Product) {
	max = products[0]
	for _, product := range products {
		if product.palindrome > max.palindrome {
			max = product
		}
	}
	return max
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	return str == reverse(str)
}

func reverse(original string) (reversed string) {
	for _, v := range original {
		reversed = string(v) + reversed
	}
	return reversed
}
