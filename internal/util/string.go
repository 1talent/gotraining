package util

import "strings"

// Lowercases a string and trims whitespace from the beginning and end of the string
func ToUsernameFormat(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}
