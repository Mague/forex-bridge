package utils

import (
	"html"
	"strings"
)

func trim(str string) string {
	return strings.TrimSpace(str)
}

func htmlScape(str string) string {
	return html.EscapeString(str)
}
func TrimHtmlScape(str string) string {
	return trim(htmlScape(str))
}
