package diffsquares

// SquareOfSum returns (1 + 2 + ... + n) ^ 2
func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares returns 1^2 + 2^2 + ... + n^2
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between the SquareOfSum(n) and SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
