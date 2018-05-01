// Package acronym has functions to give acronyms for strings
package acronym

import "strings"

// Abbreviate returns an acronym for s
func Abbreviate(s string) string {
	abb := ""
	flag := true
	for i, r := range s {
		if isAlpha(r) {
			if flag {
				abb += strings.ToUpper(s[i : i+1])
				flag = false
			}
		} else {
			flag = true
		}

	}
	return abb
}

// isAlpha checks if character is alphabet
func isAlpha(s rune) bool {
	if !(s <= 'z' && s >= 'a') && !(s <= 'Z' && s >= 'A') {
		return false
	}

	return true
}
