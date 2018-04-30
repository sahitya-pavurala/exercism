// Package luhn provides methods to implement luhn's algorithm
package luhn

import (
	"strings"
)

// Valid validates if str is a valid number
func Valid(str string) bool {
	if len(str) <= 1 {
		return false
	}

	str = strings.Replace(str, " ", "", -1)

	flag := false

	result := 0

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] < '0' || str[i] > '9' {
			return false
		}

		chr := int(str[i])
		if flag {

			val := chr * 2
			if val > 9 {
				result += val - 9
			} else {
				result += val
			}

		} else {
			result += chr
		}
		flag = !flag
	}

	if result%10 == 0 {
		return true
	}

	return false

}
