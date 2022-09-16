package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for i := 0; i < len(arguments); i++ {
		for j := 0; j < len(string(arguments[i])); j++ {
			z01.PrintRune(arguments[i][j])
		}
		z01.PrintRune('\n')
	}
}
