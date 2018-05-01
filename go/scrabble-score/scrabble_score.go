// Package scrabble has functions to give score of a word
package scrabble

import "strings"

// Score give the scrabble score of a word
func Score(word string) int {

	result := 0
	word = strings.ToUpper(word)
	for _, c := range word {
		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			result += 1
		case 'D', 'G':
			result += 2
		case 'B', 'C', 'M', 'P':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
		case 'K':
			result += 5
		case 'J', 'X':
			result += 8
		case 'Q', 'Z':
			result += 10
		}

	}

	return result
}
