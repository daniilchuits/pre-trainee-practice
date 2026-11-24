package wordcount

import (
	"strings"
	"unicode"
)

func Tokens(s string) []string {

	var (
		sb  strings.Builder
		res []string
	)

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		} else {
			if sb.Len() != 0 {
				res = append(res, sb.String())
				sb.Reset()
			}
		}
	}
	if sb.Len() != 0 {
		res = append(res, sb.String())
	}
	return res
}
