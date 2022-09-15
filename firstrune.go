package piscine

func FirstRune(s string) rune {
	if rune(s[0]) <= 122 && rune(s[0]) > 97 {
		return (rune(s[0]))
	} else {
		return (rune(s[0]) - 130)
	}
}
