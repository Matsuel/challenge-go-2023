package piscine

func IterativeFactorial(nb int) int {
	rep := 1
	for i := nb; i > 0; i-- {
		rep *= i
	}
	return rep
}
