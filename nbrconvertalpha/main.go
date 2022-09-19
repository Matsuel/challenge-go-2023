package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	// lettre := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	// corr_lettre := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26"}
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
	// var sol int

}
