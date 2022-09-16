package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	var index int
	for i := len(arguments); rune(arguments[i]) != rune(92); i-- {
		index = i
	}
	z01.PrintRune('.')
	z01.PrintRune('/')
	for i := index; i < len(arguments)-4; i++ {
		z01.PrintRune(rune(arguments[i]))
	}
	// si sa print toute la ligne c'est normal c'est windows
}
