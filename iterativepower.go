package piscine

func IterativePower(nb int, power int) int {
	rep := nb
	if power < 0 {
		return nb
	}
	for i := power; i > 1; i-- {
		rep *= nb
	}
	return rep
}
