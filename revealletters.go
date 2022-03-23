package hangmanweb

func RevealLetters(s string) string {
	var result []rune
	ruru := []rune(s)
	for i := 0; i < len(s); i++ {
		if ruru[i] != '\n' {
			if ruru[i] == 32 {
				result = append(result, 32)
			} else if ruru[i] == '-' {
				result = append(result, '-')
			} else { // remplace toutes les lettres du mot par des "_"
				result = append(result, '_')
			}
		}
	}
	var lenght int
	for _, v := range s {
		if v != 32 {
			lenght++
		}
	}
	n := (lenght / 2) - 1 // nombre de lettres à révéler dans le mot (cf consigne)
	if n <= 0 {
		n = 1
	}
	for i := 0; i < n; i++ {
		checked := false
		for !checked {
			rdm := Random(0, len(s))
			letter := ruru[rdm]
			for j, v := range result {
				if j == rdm && letter != v && letter != 32 { // si la lettre n'a pas déjà été révélée
					result[rdm] = ruru[rdm]
					checked = true
					break
				}
			}
		}
	}
	return string(result)
}
