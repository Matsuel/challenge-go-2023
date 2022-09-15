package piscine

func Join(strs []string, sep string) string {
	var rep string
	for _, el := range strs {
		rep += string(el)
		rep += sep
	}
	return rep
}
