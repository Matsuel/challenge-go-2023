package piscine

func IterativePower(nb int, power int) int {
	rep := nb
	if nb < 0 {
		return 1
	} else if power <= 0 && nb <= 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	for i := power; i > 1; i-- {
		rep *= nb
	}
	return rep
}
