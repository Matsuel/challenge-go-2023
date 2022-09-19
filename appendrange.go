package piscine

func AppendRange(min, max int) []int {
	rep := []int{}
	if min == max {
		rep = []int(nil)
	} else if min < max {
		for i := min; i < max; i++ {
			rep = append(rep, i)
		}
	}
	return rep
}
