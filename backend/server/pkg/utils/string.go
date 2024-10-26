package utils

import "regexp"

func Empty(text *string) bool {
	return text == nil || *text == ""
}

func MatchPattern(pattern string, text string) bool {
	return regexp.MustCompile(pattern).MatchString(text)
}

func IsUuid(text string) bool {
	return MatchPattern(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`, text)
}
