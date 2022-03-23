package hangmanweb

import (
	"io/ioutil"
	"log"
)

func RandomWord(dictionnary string) int {
	file, err := ioutil.ReadFile(dictionnary)
	if err != nil {
		log.Fatal(err)
	}
	var ruru []rune
	for i := 0; i < len(file); i++ {
		if file[i] != 10 {
			if file[i] == 13 {
				ruru = append(ruru, '\n')
			} else {
				ruru = append(ruru, rune(file[i]))
			}
		}
	}
	counter := 1
	for i := 0; i < len(ruru); i++ {
		if ruru[i] == '\n' {
			counter++
		}
	}
	if counter == 1 {
		return 1
	}
	return Random(1, counter)
}
