// Package bob has functions to bob's replies
package bob

import (
	"strings"
)

// Hey returns Bob's remark
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	yell := isYelling(remark)
	question := isQuestion(remark)

	if question {
		if yell {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if yell {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

// isUpper function to check if every character is in upper case
func isUpper(s string) bool {
	alphaFlag := false

	for _, r := range s {

		if r <= 'z' && r >= 'a' {
			return false
		} else if r <= 'Z' && r >= 'A' {
			alphaFlag = true
		}
	}
	if alphaFlag {
		return true
	}
	return false
}

// isYelling function to check if remark is a yelling
func isYelling(remark string) bool {
	if isUpper(remark) {
		return true
	}
	return false
}

// isQuestion function to check if remark is a question
func isQuestion(remark string) bool {
	if remark[len(remark)-1] == '?' {
		return true
	}
	return false
}
