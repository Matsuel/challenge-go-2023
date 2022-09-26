package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	rep := 0
	if len(args) == 0 {
		return
	}
	for i := 0; i < len(args); i++ {
		if i == 0 {
			if Atoi(args[0]) != 0 {
				rep = Atoi(args[0])
			} else {
				return
			}
		} else if i == 1 {
			if args[1] != "+" && args[1] != "-" && args[1] != "%" && args[1] != "/" && args[1] != "*" {
				return
			}
		} else if i == 2 {
			if args[2] == "0" && args[1] == "/" {
				err := "No division by 0"
				for _, ch := range err {
					z01.PrintRune(ch)
				}
				return
			} else if args[2] == "0" && args[1] == "%" {
				err := "No modulo by 0"
				for _, ch := range err {
					z01.PrintRune(ch)
				}
				return
			} else if args[1] == "+" {
				rep += Atoi(args[2])
			} else if args[1] == "*" {
				rep *= Atoi(args[2])
			} else if args[1] == "-" {
				rep -= Atoi(args[2])
			} else if args[1] == "/" {
				rep /= Atoi(args[2])
			}
		}
	}
	reponse := string(rep)
	for _, ch := range reponse {
		z01.PrintRune(ch + 48)
	}
}

func Atoi(s string) int {
	atoi := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			atoi = atoi*10 + int(ch-'0')
		} else {
			return 0
		}
	}
	return atoi
}
