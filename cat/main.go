package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) == 1 {
		rep := "Hello"
		for {
			for _, ch := range rep {
				z01.PrintRune(ch)
			}
			time.Sleep(800)
		}
	} else if len(arg) >= 2 {
		for i := 1; i < len(arg); i++ {
			data, err := ioutil.ReadFile(arg[i])
			if err == nil {
				for _, ch := range string(data) {
					z01.PrintRune(ch)
				}
			} else if err != nil && i == 1 {
				rep := "ERROR: open " + string(arg[1]) + ": No such file or directory"
				for _, ch := range rep {
					z01.PrintRune(ch)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			} else if err != nil && i != 1 {
				rep := "ERROR: " + string(arg[1]) + ": No such file or directory"
				for _, ch := range rep {
					z01.PrintRune(ch)
				}
				z01.PrintRune('\n')
			}
		}
	}
}
