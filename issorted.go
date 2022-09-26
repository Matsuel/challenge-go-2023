package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if f(a[0], a[1]) > 0 {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
		return true
	} else if f(a[0], a[1]) < 0 {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
		return true
	}
}
