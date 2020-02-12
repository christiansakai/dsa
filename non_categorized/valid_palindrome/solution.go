package solution

import (
	"regexp"
	"strings"
)

func Solve(s string) bool {
	if len(s) <= 1 {
		return true
	}

	s = strings.ToLower(s)

	start := 0
	end := len(s) - 1
	for {
		for ; start < len(s) && !isAlphaNumeric(s[start]); start++ {
		}
		for ; end >= 0 && !isAlphaNumeric(s[end]); end-- {
		}

		if start >= end {
			break
		}

		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func isAlphaNumeric(c byte) bool {
	match, _ := regexp.Match("[a-zA-Z0-9]+", []byte{c})
	return match
}
