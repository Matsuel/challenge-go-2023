package piscine

func SplitWhiteSpaces(s string) []string {
	rep := []string{}
	mot := ""
	for _, ch := range s {
		if string(ch) == " " {
			rep = append(rep, mot)
			mot = ""
		} else {
			mot += string(ch)
		}
	}
	rep = append(rep, mot)
	return rep
}
