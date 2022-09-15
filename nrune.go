package piscine

func NRune(s string, n int) rune {
	if n <= len(s) {
		return []rune(s)[n-1]
	}else{
		return 0
	}
}
