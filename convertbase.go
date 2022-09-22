package piscine

func ConvertBase(nbr, baseFrom, baseTo string) int {
	coef := 1
	rep := 0
	t := len(nbr) - 1
	if baseFrom == "01" {
		for i := 0; i < len(nbr); i++ {
			coef = 1
			if string(nbr[i]) == "1" {
				for j := 0; j < t; j++ {
					coef *= 2
				}
				rep += coef
				t -= 1
			} else {
				t -= 1
			}
		}
	}
	return rep
}
