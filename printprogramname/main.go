package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	for _, el := range arguments {
		z01.PrintRune(el)
	}
	// var elem string
	// for _, el := range arguments {
	// 	if el == rune(92) {
	// 		elem = ""
	// 	} else {
	// 		elem += string(el)
	// 	}
	// }
	// for i := 0; (rune(elem[i]) != rune(46)) || i < len(elem)-1; i++ {
	// 	z01.PrintRune(rune(elem[i]))
	// }
}
