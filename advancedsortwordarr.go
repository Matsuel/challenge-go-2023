package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == 1 {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
}
