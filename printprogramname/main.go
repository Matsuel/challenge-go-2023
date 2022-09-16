package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for _, el := range arguments {
		if (el != rune(92)) && (el != rune(46)) && (el != rune(47)) {
			z01.PrintRune(el)
		}
	}
	// si sa print toute la ligne c'est normal c'est windows
}
