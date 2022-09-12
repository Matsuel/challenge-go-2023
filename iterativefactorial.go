package piscine

func IterativeFactorial(nb int) int {
	rep := 1
	if nb <= 0 {
		return 0
	} else if nb > 0 {
		for i := nb; i > 0; i-- {
			rep *= i
		}
	}
	return rep
}
