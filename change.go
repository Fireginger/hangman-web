package hangmanweb

func Change(entry string, wordToComplete string, Word string) string {
	result := ""
	lettre := rune(entry[0])
	for i := 0; i < len(wordToComplete); i++ {
		if rune(Word[i]) == lettre {
			result += string(Word[i])
			continue
		}
		result += string(wordToComplete[i])
	}
	return result
}
