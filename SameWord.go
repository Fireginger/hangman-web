package hangmanweb

func SameWord(str string, strWithSpace string) bool {
	if len(str) <= 0 {
		return false
	}
	ruru := []rune(str)
	var ruruNoSpace []rune
	for _, v := range strWithSpace {
		if v != ' ' {
			ruruNoSpace = append(ruruNoSpace, v)
		}
	}
	for i := 0; i < len(ruruNoSpace); i++ {
		if ruruNoSpace[i] != ruru[i] {
			return false
		}
	}
	return true
}
