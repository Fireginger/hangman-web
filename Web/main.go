package main

import (
	"fmt"
	"hangmanweb"
	"net/http"
	"text/template"
)

type Page struct {
	Title       *string
	Sub         string
	StateLetter *string
}

var (
	counter  int
	Jeu      bool = false
	Victoire bool = false
	Accueil  bool = true
	tmpl     *template.Template
)

func main() {
	tmpl, _ = template.ParseGlob("./Web/templates/*.html")
	// Set routing rules
	http.HandleFunc("/", FunctionManager)

	fileserver := http.FileServer(http.Dir("./Web/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileserver))
	//Use the default DefaultServeMux.
	http.ListenAndServe(":8080", nil)

}

func FunctionManager(w http.ResponseWriter, r *http.Request) {
	if Accueil {
		Welcome(w, r)
	}
	if !Jeu && !Accueil {
		Home(w, r)
		fmt.Println("1")
	}
	if Jeu && !Victoire {
		GameMode(w, r)
		fmt.Println("2")
	}
	if Victoire {
		EndMode(w, r)
		fmt.Println("3")
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	AllMyMot := r.Form["PLAY"]
	if len(AllMyMot) > 0 {
		MyButton := AllMyMot[0]
		fmt.Println(MyButton)

		if MyButton == "PLAY" {
			Accueil = false
		}

	}

	if Accueil {
		tmpl.ExecuteTemplate(w, "welcome", tmpl)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Get Getted")
	}
	MainVarReset()
	r.ParseForm()
	AllMyMot := r.Form["subject"]
	//var MyMot string
	if len(AllMyMot) > 0 {
		MyButton := AllMyMot[0]
		fmt.Println(MyButton)

		if MyButton == "EASY" {
			hangmanweb.HangmanFileInit(1)
		} else if MyButton == "MEDIUM" {
			hangmanweb.HangmanFileInit(2)
		} else if MyButton == "HARD" {
			hangmanweb.HangmanFileInit(3)
		}
		fmt.Println(hangmanweb.TempFile)
		hangmanweb.HangmanWebInit()
		Jeu = true
	}
	if !Jeu {
		fmt.Println("INGAME")
		tmpl.ExecuteTemplate(w, "home", "")
	}

}

func GameMode(w http.ResponseWriter, r *http.Request) {

	counter++

	/*data := Page{
		Title: "Wesh Wesh",
		Sub:   "Canne à peche",
	}
	tmpl.ExecuteTemplate(w, "index", data)*/

	MyGame := Page{
		Title:       &hangmanweb.CompleteWord,
		Sub:         " ",
		StateLetter: &hangmanweb.LetterState,
	}
	r.ParseForm()
	gobackbutton := r.Form["GOBACK"]
	if len(gobackbutton) > 0 {
		GoBack := gobackbutton[0]
		if GoBack == "GOBACK" {
			fmt.Println("GOBACK CLICKED")
			Jeu = false
			Home(w, r)
		}
	}

	AllMyMot := r.Form["MyLetter"]
	if len(AllMyMot) > 0 {
		MyMot := AllMyMot[0]
		if len(MyMot) > 0 {
			MyGame.StateLetter = hangmanweb.HangmanWebPlay(MyMot)
		}
		fmt.Println(hangmanweb.LetterState)
	}
	MyGame.Sub = "T'as encore " + hangmanweb.MyAttemptToString(&hangmanweb.Attempts) + " vies"
	//WinString := "YOU WIN!"
	if hangmanweb.Complete || hangmanweb.Attempts <= 0 {
		Victoire = true
		//MyGame.Title = &WinString
	}
	if Jeu && !Victoire {
		tmpl.ExecuteTemplate(w, "index", MyGame)
		//tmpl.ExecuteTemplate(w, "lettreused", "")
	}
	//http.ServeFile(w, r, "./Web/templates/index.html")
	fmt.Println(hangmanweb.Word + " est actuellement " + hangmanweb.CompleteWord)
}

func EndMode(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	AllMyMot := r.Form["PLAYAGAIN"]

	if len(AllMyMot) > 0 {
		MyButton := AllMyMot[0]
		if MyButton == "PLAYAGAIN" {
			fmt.Println("PlayAgain Clicked")
			MainVarReset()
		} else {
			fmt.Println("BACK TO HOME")
			fmt.Println("BACKTOHOME Clicked")
			Accueil = true
			MainVarReset()
		}
	}
	WinOrLose := Page{
		Title:       &hangmanweb.Word,
		Sub:         "",
		StateLetter: &hangmanweb.LetterState,
	}
	if hangmanweb.Complete {
		WinOrLose.Sub = " You WIN! The Word was " + hangmanweb.Word
	} else {
		WinOrLose.Sub = " You LOSE! The Word was " + hangmanweb.Word
	}
	if Victoire {
		tmpl.ExecuteTemplate(w, "endscreen", WinOrLose)
		fmt.Println("Winscreen")
	} else if !Accueil {
		fmt.Println("PLAY AGAIN")
		Home(w, r)
	} else {
		fmt.Println("Back")
		Welcome(w, r)
	}

}

func MainVarReset() {
	Jeu = false
	hangmanweb.Complete = false
	hangmanweb.Attempts = 10
	Victoire = false
	fmt.Println("RESET VARIABLE")
	var temp []rune
	hangmanweb.AllUsedLetters = temp
}
