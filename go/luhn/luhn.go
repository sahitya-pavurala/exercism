// Package luhn provides methods to implement luhn's algorithm
package luhn

import (
	"strconv"
	"strings"
)

// Valid validates if str is a valid number
func Valid(str string) bool {
	if len(str) <= 1 {
		return false
	}

	str = strings.Replace(str, " ", "", -1)

	flag := false

	str2 := ""

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] < '0' || str[i] > '9' {
			return false
		}

		if flag {

			val := (int)(str[i]-48) * 2
			if val > 9 {
				str2 = strconv.Itoa(val-9) + str2
			} else {
				str2 = strconv.Itoa(val) + str2
			}

		} else {
			str2 = string(str[i]) + str2
		}
		flag = !flag
	}

	result := 0
	for _, c := range str2 {
		result += int(c) - 48
	}

	if result%10 == 0 {
		return true
	}

	return false

}
