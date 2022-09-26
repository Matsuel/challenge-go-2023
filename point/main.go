package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	list := []string{}
	list = append(list, string(points.x))
	list = append(list, string(points.y))
	z01.PrintRune('x')
	z01.PrintRune(rune(32))
	z01.PrintRune('=')
	z01.PrintRune(rune(32))
	for _, ch := range list[0] {
		z01.PrintRune(ch)
	}
	z01.PrintRune('y')
	z01.PrintRune(rune(32))
	z01.PrintRune('=')
	z01.PrintRune(rune(32))
	for _, ch := range list[1] {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
