package util

import (
	"strings"
	"regexp"
)

func StringToPath(s string) string {

	s = strings.ToLower(s)

	respacing := regexp.MustCompile(`[ ]+`)
	rechars := regexp.MustCompile(`[^ a-z0-9]`)

	s = respacing.ReplaceAllString(s, " ")
	s = rechars.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "-")

	return s

}