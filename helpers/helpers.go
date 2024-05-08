package helpers

import (
	"regexp"
	"strconv"
	"time"
)

func Parse(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func ValidateString(s string) bool {
	if s == "" {
		return false
	}
	nameAuthorRegex := regexp.MustCompile(`^[a-zA-Z0-9\s.,'-]+$`)
	return nameAuthorRegex.MatchString(s)
}

func ValidateNumber(a int) bool {
	return a <= time.Now().Year() && a > 999
}
