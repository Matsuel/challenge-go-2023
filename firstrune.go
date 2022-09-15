package piscine

func FirstRune(s string) rune {
	if rune(s[0]) <= 122 && rune(s[0]) > 97 || rune(s[0]) <= 57 && rune(s[0]) > 47 || rune(s[0]) <= 90 && rune(s[0]) > 64 {
		return (rune(s[0]))
	} else {
		return (rune(s[0]) - 130)
	}
}
