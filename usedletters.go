package hangmanweb

import "fmt"

func UsedLetters(entry string, ruru *[]rune) {
	if len(*ruru) != 0 {
		*ruru = append(*ruru, '-')
	}
	for _, v := range entry {
		*ruru = append(*ruru, v)
	}
	if len(*ruru) != 0 {
		fmt.Print("Answers given : " + string(*ruru) + "\n")
	}
}
