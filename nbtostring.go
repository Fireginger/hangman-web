package hangmanweb

func NbToString(n int) string {
	var arr []rune
	var result string
	nbr := n
	for {
		if nbr >= 1 || nbr <= -1 {
			var digit int
			if n > 0 {
				digit = nbr % 10
				nbr -= digit
			} else {
				digit = 0 - (nbr % 10)
				nbr += digit
			}
			arr = append(arr, rune(digit+48))
			nbr /= 10
		} else {
			break
		}
	}
	if n == 0 {
		result = "0"
	} else {
		for j := len(arr) - 1; j >= 0; j-- {
			result += string(arr[j])
		}
	}
	return result
}

func MyAttemptToString(nn *int) string {
	n := *nn
	var arr []rune
	var result string
	nbr := n
	for {
		if nbr >= 1 || nbr <= -1 {
			var digit int
			if n > 0 {
				digit = nbr % 10
				nbr -= digit
			} else {
				digit = 0 - (nbr % 10)
				nbr += digit
			}
			arr = append(arr, rune(digit+48))
			nbr /= 10
		} else {
			break
		}
	}
	if n == 0 {
		result = "0"
	} else {
		for j := len(arr) - 1; j >= 0; j-- {
			result += string(arr[j])
		}
	}
	return result
}
