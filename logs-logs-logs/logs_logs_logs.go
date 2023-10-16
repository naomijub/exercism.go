package logs

import (
	"strings"
	"unicode/utf8"
)

const (
	Default        = "default"
	Recommendation = "recommendation"
	Search         = "search"
	Weather        = "weather"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, rune := range log {
		switch rune {
		case '❗':
			{
				return Recommendation
			}
		case '🔍':
			{
				return Search
			}
		case '☀':
			{
				return Weather
			}
		}
	}
	return Default
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var sb strings.Builder

	for _, aRune := range log {
		if aRune == oldRune {
			sb.WriteRune(newRune)
		} else {
			sb.WriteRune(aRune)
		}
	}

	return sb.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	len := utf8.RuneCountInString(log)
	return len <= limit
}
