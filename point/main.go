package main

import "github.com/01-edu/z01"

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
	pt := ""
	pt += "x = "
	for _, ch := range string(points.x) {
		pt += string(ch)
	}
	pt += ", y = "
	for _, ch := range string(points.y) {
		pt += string(ch)
	}
	for i := 0; i < len(pt); i++ {
		z01.PrintRune(rune(pt[i]))
	}
	z01.PrintRune('\n')
}
