package utils

import "strings"

// Transforms a given string into one with all uppercase letters and an exclamation point.
func MakeExcited(s string) string {
	excited := strings.ToUpper(s) + "!"
	return excited
}
