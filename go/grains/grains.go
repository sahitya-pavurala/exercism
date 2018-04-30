// Package grains provides methods to return the number of grain in a chessboard
package grains

import (
	"fmt"
	"math"
)

// Square returns number of grains in a chess board square
func Square(sqNum int) (uint64, error) {

	if sqNum <= 0 || sqNum > 64 {
		return 0, &illegalArgError{"Illegal argument provided, numbr should be greater than 0"}
	}

	result := math.Pow(float64(2), float64(sqNum-1))

	return uint64(result), nil

}

// Total returns total number of grains on the chess board
func Total() uint64 {
	var result uint64

	for i := 0; i < 64; i++ {
		result += uint64(math.Pow(float64(2), float64(i)))
	}

	return result
}

// // illegalArgError is struct for custom error type
type illegalArgError struct {
	message string
}

// illegalArgErrors implementing Error
func (error *illegalArgError) Error() string {
	return fmt.Sprintf("%s", error.message)

}
