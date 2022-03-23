package hangmanweb

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Picture(model string, n int, rows int) {
	file, err := ioutil.ReadFile(model)
	if err != nil {
		log.Fatal(err)
	}
	str := ""
	counter := 0

	start := 0
	end := len(file) - 1
	for i := 0; i < len(file); i++ {
		if counter == (n-1)*rows {
			start = i - (rows + 1)
			if start < 0 {
				start = 0
			}
		}
		if counter == n*rows {
			end = i - (rows + 1)
			if end < 0 {
				end = 0
			}
		}
		if file[i] == 13 {
			counter++
		}
	}
	if n != 10 {
		str = string(file[start:end])
	} else {
		str = string(file[start:])
	}
	fmt.Print(str)
}
