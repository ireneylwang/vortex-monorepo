package utils

func ValidateMissingField(text string) bool {
	return Empty(&text)
}

func ValidateInvalidFormat(format string, text string) bool {
	return !MatchPattern(format, text)
}
