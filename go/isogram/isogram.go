// Package isogram provides functions to check is string is isogram
package isogram

import "strings"

// IsIsogram returns true if word is a isogram, false otherwise
func IsIsogram(word string) bool {
	checker := make(map[rune]bool)

	word = strings.ToLower(word)

	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			if _, ok := checker[c]; ok {
				return false
			}
			checker[c] = true
		}
	}
	return true
}
