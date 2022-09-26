package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	if arg[0]== "--insert" ||  arg[0]== "-i"{
		fmt.Println("This flag inserts the string into the string passed as argument.")
	}else if arg[0]== "--order" || arg[0]== "-o"{
		fmt.Println("This flag will behave like a boolean, if it is called it will order the argument.")
	}
	for arguments := range arg {
		if string(arguments) == "--insert" || string(arguments) == "-i" {
			fmt.Println("")
		}
	}
}
