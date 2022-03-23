package hangmanweb

func IsComplete(Word string, str string) bool {
	ruruToFind := []rune(Word)
	ruru := []rune(str)
	for i := 0; i < len(str); i++ {
		if ruruToFind[i] != ruru[i] {
			return false
		}
	}
	return true
}
