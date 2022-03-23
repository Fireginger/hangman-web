package hangmanweb

func ToUpper(str string) string {
	result := ""
	ruru := []rune(str)
	for i := 0; i < len(ruru); i++ {
		if ruru[i] >= 97 && ruru[i] <= 122 {
			result += string(ruru[i] - 32)
			continue
		}
		result += string(str[i])
	}
	return result
}
