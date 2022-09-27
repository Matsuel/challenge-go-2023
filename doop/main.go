package main

import (
	"os"
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
				os.Stdout.WriteString("No division by 0")
				return
			} else if args[2] == "0" && args[1] == "%" {
				os.Stdout.WriteString("No modulo by 0")
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
	res := []string{}
	for i := 0; rep > 0; i++ {
		a := rep % 10
		res = append(res, string(a+48))
		rep -= a
		rep /= 10
	}
	for i := len(res) - 1; i >= 0; i-- {
		os.Stdout.WriteString(res[i])
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