package piscine

func ToUpper(s string) string {
	var rep string
	for _, ch := range s {
		if IsLower(string(ch)) {
			rep += string(rune(ch - 22))
		} else {
			rep += string(ch)
		}
	}
	return rep
}
