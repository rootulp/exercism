package palindrome

import "strconv"

// Define Product type here.
type Product struct {
	palindrome     int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (min Product, max Product, e error) {
	products, e := getPalindromeProducts(fmin, fmax)
	if e != nil {
		return min, max, e
	}
	min = getMin(products)
	max = getMax(products)
	return min, max, nil
}

func getPalindromeProducts(min int, max int) (products []Product, e error) {
	productToFactors := map[int][][2]int{}
	products = make([]Product, 0)
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			candidate := i * j
			if isPalindrome(candidate) {
				factor := [2]int{i, j}
				factors, ok := productToFactors[candidate]
				if ok {
					factors = append(factors, factor)
					productToFactors[candidate] = factors
				} else {
					productToFactors[candidate] = [][2]int{factor}
				}
			}
		}
	}
	for palindrome, factors := range productToFactors {
		product := Product{
			palindrome:     palindrome,
			Factorizations: factors,
		}
		products = append(products, product)
	}
	return products, nil
}

func getMin(products []Product) (min Product) {
	if len(products) == 0 {
		return Product{}
	}
	min = products[0]
	for _, product := range products {
		if product.palindrome < min.palindrome {
			min = product
		}
	}
	return min
}

func getMax(products []Product) (max Product) {
	if len(products) == 0 {
		return Product{}
	}
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
