package utils

import "regexp"

var phoneRegex = regexp.MustCompile(`^\+?1?\d{10}$`)

func IsValidPhone(s string) bool {
	return phoneRegex.MatchString(s)
}