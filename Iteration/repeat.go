package iteration

import "strings"

var repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
