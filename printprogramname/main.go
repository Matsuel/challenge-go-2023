package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	i := len(arguments)
	for rune(arguments[i]) != rune(92) {
		i--
	}
	z01.PrintRune('.')
	z01.PrintRune('/')
	for i < len(arguments)-4 {
		z01.PrintRune(rune(arguments[i]))
		i++
	}
	// si sa print toute la ligne c'est normal c'est windows
}
