package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	if arguments[1] == "--upper" {
		for i := 2; i < len(arguments); i++ {
			a := piscine.Atoi(arguments[i])
			if a == 0 {
				z01.PrintRune(rune(32))
			} else if a <= 26 {
				z01.PrintRune(rune(a + 64))
			} else {
				z01.PrintRune(rune(32))
			}
		}
	} else {
		for i := 1; i < len(arguments); i++ {
			a := piscine.Atoi(arguments[i])
			if a == 0 {
				z01.PrintRune(rune(32))
			} else if a <= 26 {
				z01.PrintRune(rune(a + 96))
			} else {
				z01.PrintRune(rune(32))
			}
		}
	}
}
