package iteration

func Repeat(character string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += character
	} 
	return result
}

