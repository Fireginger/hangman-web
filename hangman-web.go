package hangmanweb

import (
	"fmt"
)

// On définit les variables
var (
	file           string
	LetterState    string
	TempFile       string
	Word           string = ""
	CompleteWord   string = ""
	AllUsedLetters []rune
	Attempts       int     = 10
	Complete       bool    = false
	ptrUsed        *[]rune = &AllUsedLetters
)

func HangmanFileInit(nb int) {
	if nb == 1 {
		file = "words.txt"
		TempFile = file
		return
	}

	file = "words" + NbToString(nb) + ".txt"
	TempFile = file
}

func HangmanWebInit() {
	// CHARGEMENT DES VARIABLE
	// prendre le WordToFind
	nb := RandomWord(file)
	Word = GetWord(file, nb)
	Word = ToUpper(Word)
	// detecter les input et les valider
	CompleteWord = RevealLetters(Word)
	fmt.Println("Good Luck, you have " + NbToString(Attempts) + " Attempts.")
	println(CompleteWord)
	fmt.Println()
}

func HangmanWebPlay(entry string) *string {

	// CE QUI SE PASSE IN GAME
	var HavetoReturn bool
	var RetrunState string
	entry, HavetoReturn, RetrunState = TreatInput(entry, Word, ptrUsed)
	if HavetoReturn {
		return &RetrunState
	}
	entry = ToUpper(entry)
	var letter rune
	for _, v := range entry {
		letter = v
		break
	}

	if len(entry) == 1 {
		if IsRight(Word, letter) {
			oldstr := CompleteWord
			CompleteWord = Change(entry, CompleteWord, Word)
			if oldstr != CompleteWord {
				Picture("thumb.txt", 1, 15)
				UsedLetters(entry, ptrUsed)
				println(CompleteWord)
				fmt.Println()
				Complete = IsComplete(CompleteWord, Word)

			} else {
				Attempts--
				if Attempts > 0 {
					fmt.Println("No more of this letter in the word, " + NbToString(Attempts) + " attempt(s) remaining\n")
					RetrunState = "BAD CHOISE"
					Picture("hangman.txt", 10-Attempts, 8)
					UsedLetters(entry, ptrUsed)
					println(CompleteWord)
					fmt.Println()
				}
			}

		} else {
			Attempts--
			if Attempts > 0 {
				fmt.Println("Not present in the word, " + NbToString(Attempts) + " attempt(s) remaining\n")
				RetrunState = "BAD CHOISE"
				Picture("hangman.txt", 10-Attempts, 8)
				UsedLetters(entry, ptrUsed)
				println(CompleteWord)
				fmt.Println()
			}
		}
	} else {
		if SameWord(entry, Word) {
			CompleteWord = entry
			Picture("thumb.txt", 1, 15)
			println(RemoveSpace(Word))
			fmt.Println()
			Complete = true

		} else {
			Attempts -= 2
			if Attempts > 0 {
				fmt.Println("Incorrect word, " + NbToString(Attempts) + " attempt(s) remaining\n")
				RetrunState = "BAD CHOISE"
				Picture("hangman.txt", 10-Attempts, 8)
				UsedLetters(entry, ptrUsed)
				println(CompleteWord)
			}
		}
	}

	if Complete {
		fmt.Println("Congrats !")
	} else if Attempts <= 0 {
		Picture("hangman.txt", 10, 8)
		fmt.Println("You failed to save José. He died.")
		fmt.Println("The word was : ")
		println(RemoveSpace(Word))
		fmt.Println()

	}

	return &RetrunState
}
