package piscine

func IterativePower(nb int, power int) int {
	rep := nb
	for i := power; i > 1; i-- {
		rep *= nb
	}
	return rep
}
