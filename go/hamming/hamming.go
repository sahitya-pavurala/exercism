//Package hamming has methods to get the hamming ditance betwenn two DNA signatures
package hamming

import "fmt"

// Distance returns the hamming distance between two strands
func Distance(a, b string) (int, error) {

	result := 0

	if len(a) != len(b) {
		return -1, &illegalArgError{"Illegal arguments provided, Required sequences of same length"}
	}

	for idx := range a {
		if a[idx] != b[idx] {
			result += 1
		}

	}

	return result, nil

}

// illegalArgError is struct for custom error type
type illegalArgError struct {
	message string
}

// illegalArgErrors implementing Error
func (error *illegalArgError) Error() string {

	return fmt.Sprintf("%s", error.message)
}
