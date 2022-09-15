package piscine

func IsLOwer(s string) bool {
	for _, ch := range s {
		if rune(ch) < 97 || rune(ch) > 122 {
			return false
		}
	}
	return true
}
