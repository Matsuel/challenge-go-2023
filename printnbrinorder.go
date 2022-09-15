package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	n_tab := []rune{}
	if n <= 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			div := n % 10
			n /= 10
			n_tab = append(n_tab, rune(div+48))
		}
		for i := 0; i < len(n_tab); i++ {
			z01.PrintRune(n_tab[i])
		}
	}
}
