package iteration

import "strings"

func Repeat(character string, count int) string {
	result := strings.Repeat(character, count)	
	return result
}
