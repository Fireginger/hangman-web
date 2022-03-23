package hangmanweb

import (
	"io/ioutil"
	"log"
)

func GetWord(dictionnary string, nb int) string {
	file, err := ioutil.ReadFile(dictionnary)
	if err != nil {
		log.Fatal(err)
	}
	remove := RemoveAccent(file)
	counter := 1
	start := 0
	str := ""
	for i := 0; i < len(remove); i++ {
		if remove[i] == '\n' {
			counter++
		}
		if counter == nb {
			for j := i + start; j < len(remove); j++ {
				if remove[j] == '\n' {
					break
				}
				str += string(remove[j])
				if j < len(remove)-1 {
					if remove[j+1] != '\n' {
						str += " "
					}
				}
			}
			break
		}
		start = 1
	}
	return str
}
