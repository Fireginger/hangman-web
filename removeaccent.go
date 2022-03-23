package hangmanweb

func RemoveAccent(byby []byte) []rune {
	var ruru []rune

	for i := 0; i < len(byby); i++ {
		if byby[i] == 13 {
			ruru = append(ruru, '\n')
		} else if (byby[i] >= 160 && byby[i] <= 164) || (byby[i] >= 128 && byby[i] <= 132) {
			ruru = append(ruru, 'A')
		} else if (byby[i] >= 168 && byby[i] <= 171) || (byby[i] >= 136 && byby[i] <= 139) {
			ruru = append(ruru, 'E')
		} else if (byby[i] >= 172 && byby[i] <= 175) || (byby[i] >= 140 && byby[i] <= 143) {
			ruru = append(ruru, 'I')
		} else if (byby[i] >= 179 && byby[i] <= 182) || (byby[i] >= 146 && byby[i] <= 150) {
			ruru = append(ruru, 'O')
		} else if (byby[i] >= 185 && byby[i] <= 188) || (byby[i] >= 153 && byby[i] <= 156) {
			ruru = append(ruru, 'U')
		} else if byby[i] == 191 || byby[i] == 159 {
			ruru = append(ruru, 'Y')
		} else if byby[i] == 167 || byby[i] == 135 {
			ruru = append(ruru, 'C')
		} else if byby[i] == 177 || byby[i] == 145 {
			ruru = append(ruru, 'N')
		} else if (byby[i] >= 97 && byby[i] <= 122) || (byby[i] >= 65 && byby[i] <= 90) || byby[i] == 32 || byby[i] == '-' {
			ruru = append(ruru, rune(byby[i]))
		}
	}
	return ruru
}
