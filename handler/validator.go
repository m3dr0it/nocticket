package handler

import (
	"regexp"
)

func isEmailValid(email string) bool {
	if len(email) < 1 {
		return false
	}

	if !isValidEmail(email) {
		return false
	}

	return true
}

func isValidEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
