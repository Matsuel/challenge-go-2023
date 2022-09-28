package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Printf("No enough arguments")
		os.Exit(1)
	} else {
		if args[0] == "-c" {
			data, err := ioutil
		}
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

func Itoa(i int) string {
	itoa := ""
	for i != 0 {
		ch := i % 10
		i /= 10
		itoa += string(ch + '0')
	}
	res := ""
	for i := len(itoa) - 1; i >= 0; i-- {
		res += string(itoa[i])
	}
	return res
}
