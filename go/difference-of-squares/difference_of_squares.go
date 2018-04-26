// Package diffsquares provides methods to give difference between square of the result and the result of the squares
package diffsquares

// SquareOfSums return square of the sum of n numbers
func SquareOfSums(n int) int {
	result := n * (n + 1) / 2

	return result * result
}

//SumOfSquares returns sum of the squares of n numbers
func SumOfSquares(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		result += i * i
	}

	return result
}

// Difference return the difference between SquareOfSums and SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)

}
