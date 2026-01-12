package utils

import "regexp"

var phoneRegex = regexp.MustCompile(`^\+?1?\d{10}$`)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidPhone(s string) bool {
	return phoneRegex.MatchString(s)
}


func IsValidEmail(s string) bool {
	return emailRegex.MatchString(s)
}

