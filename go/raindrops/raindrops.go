// Package raindrops has methods to convert numbers to strings
package raindrops

import "fmt"

// Convert return a string after converting num
func Convert(num int) string {
	result := ""

	if num%3 == 0 {
		result += "Pling"
	}

	if num%5 == 0 {
		result += "Plang"
	}

	if num%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = fmt.Sprint("", num)
	}

	return result
}
