package hangmanweb

import (
	"math/rand"
	"time"
)

func Random(min int, max int) int {
	if min == max || min > max {
		return min
	}
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
