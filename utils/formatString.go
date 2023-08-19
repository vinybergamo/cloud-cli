package utils

import "strings"

func FormatString(string string) string {
	lowerCase := strings.ToLower(string)

	return strings.ReplaceAll(lowerCase, " ", "-")
}
