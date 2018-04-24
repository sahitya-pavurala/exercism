// Package twofer should have a package comment that summarizes what it's about.
package twofer

import "fmt"

// ShareWith returns a string based on s
func ShareWith(s string) string {

	result := fmt.Sprintf("One for %s, one for me.", s)
	if s == "" {
		result = "One for you, one for me."
	}
	return result
}
