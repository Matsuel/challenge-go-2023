package piscine

func RecursiveFactorial(nb int) int {
	if nb == 4 {
		return 24
	} else if nb == 1 {
		return nb
	} else {
		return (nb * RecursiveFactorial(nb-1))
	}
}
