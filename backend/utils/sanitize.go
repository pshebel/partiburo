package utils

import (
	"strings"
)

func SanitizeEmail(email string) string {
	// 1. Convert to lowercase
	// 2. Trim leading/trailing whitespace, tabs, and newlines
	// 3. strings.Fields + Join handles internal weird whitespace (optional)
	
	cleanEmail := strings.ToLower(strings.TrimSpace(email))
	
	return cleanEmail
}