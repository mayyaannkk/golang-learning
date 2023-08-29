package iteration

func Repeat(character string, count int) string {
	var output string
	for i := 0; i < count; i++ {
		output += character
	}
	return output
}
