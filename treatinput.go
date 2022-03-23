package hangmanweb

import (
	"fmt"
)

func TreatInput(entry string, word string, previousEntries *[]rune) (string, bool, string) {
	treatedStr := string(RemoveAccent([]byte(entry)))
	treatedStr = ToUpper(treatedStr)
	haveToCheck := false
	if (len(treatedStr) == 1 || len(treatedStr) == len(RemoveSpace(word))) && entry != "-" {
		alreadySubmitted := false
		usedLetters := *previousEntries
		for index, v := range usedLetters {
			haveToCheck = false
			if index >= 1 && index < len(usedLetters)-1 && usedLetters[index-1] == '-' && usedLetters[index+1] == '-' {
				haveToCheck = true
				fmt.Println("Check 1")
			} else if index == 0 {
				if len(usedLetters) > 0 {
					if index < len(usedLetters)-1 && usedLetters[index+1] == '-' {
						haveToCheck = true
						fmt.Println("Check 2")

					}
				} else {
					haveToCheck = true
					fmt.Println("Check 3")

				}
			} else if index == len(usedLetters)-1 && usedLetters[index-1] == '-' {
				haveToCheck = true
				fmt.Println("Check 4")
			}
			if haveToCheck {
				if string(v) == treatedStr {
					alreadySubmitted = true
					fmt.Println("Check Action")

				}
			}
		}
		if alreadySubmitted {
			fmt.Println("You already tried this letter")
			return "", true, "You already tried this letter"
		}
		return treatedStr, false, "NICE CHOISE!"
	} else {
		fmt.Println("Invalid input, you have to write either one letter or the whole word!")
		return "", true, "Invalid input, you have to write either one letter or the whole word!"
	}
}

func RemoveSpace(str string) string {
	var result string
	for _, v := range str {
		if v != 32 {
			result += string(v)
		}
	}
	return result
}

func InputStop(str string) bool {
	strUpper := ToUpper(str)
	treatedStr := []rune(strUpper)
	saveStr := "STOP"
	if strUpper != saveStr {
		return false
	}
	for i := 0; i < len(treatedStr); i++ {
		if treatedStr[i] != rune(saveStr[i]) {
			return false
		}
	}
	return true
}
