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
	p_y := string(points.y)
	p_x := string(points.x)
	z01.PrintRune('x')
	z01.PrintRune(rune(32))
	z01.PrintRune('=')
	z01.PrintRune(rune(32))
	for _, ch := range p_x {
		z01.PrintRune(ch + 48)
	}
	z01.PrintRune('y')
	z01.PrintRune(rune(32))
	z01.PrintRune('=')
	z01.PrintRune(rune(32))
	for _, ch := range p_y {
		z01.PrintRune(ch + 48)
	}
	z01.PrintRune('\n')

}
