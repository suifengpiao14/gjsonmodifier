package gsjsonmodifier

import "strings"

func TrimSpaces(s string) string {
	return strings.Trim(s, "\r\n\t\v\f ")
}
