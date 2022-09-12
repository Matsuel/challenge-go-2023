package piscine

func IterativeFactorial(nb int) int {
	rep := 1
	if nb < 0 {
		return 0
	} else if nb > 10 || nb == 1 {
		return 1
	} else if nb > 0 {
		for i := nb; i > 0 && i < 10; i-- {
			rep *= i
		}
	}
	return rep
}
