package iteration

import "strings"

// Reapeats the string by the number of count
func Repeat(character string, count int32) string {
	var repeated strings.Builder

	for count > 0 {
		repeated.WriteString(character)
		count--
	}
	return repeated.String()
}
