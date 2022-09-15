package piscine

func Capitalize(s string) string {
	rep := ""
	for i := 0; i < len(s); i++ {
		if i == 0 && IsLower(string(s[i])) {
			b := rune(s[i] - 32)
			rep += string(b)
		} else if IsLower(string(s[i])) && (!IsAlpha(string(s[i-1]))) {
			b := rune(s[i] - 32)
			rep += string(b)
		} else {
			rep += string(s[i])
		}
	}
	return rep
}
