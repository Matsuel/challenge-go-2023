// package main

// import "piscine"

// func main() {
// 	piscine.IsNegative(1)
// 	piscine.IsNegative(0)
// 	piscine.IsNegative(-1)
// }

// func main() {
// 	piscine.PrintComb()
// }

// func main() {
// 	piscine.PrintComb2()
// }

// package main

// import (
// 	"piscine"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	piscine.PrintNbr(-123)
// 	piscine.PrintNbr(0)
// 	piscine.PrintNbr(123)
// 	z01.PrintRune('\n')
// }

// package main

// import "piscine"

// func main() {
// 	piscine.PrintCombN(1)
// 	piscine.PrintCombN(3)
// 	piscine.PrintCombN(9)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	n := 0
// 	piscine.PointOne(&n)
// 	fmt.Println(n)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 0
// 	b := &a
// 	n := &b
// 	piscine.UltimatePointOne(&n)
// 	fmt.Println(a)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 13
// 	b := 2
// 	var div int
// 	var mod int
// 	piscine.DivMod(a, b, &div, &mod)
// 	fmt.Println(div)
// 	fmt.Println(mod)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 13
// 	b := 2
// 	piscine.UltimateDivMod(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// package main

// import "piscine"

// func main() {
// 	piscine.PrintStr("Hello World!")
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 0
// 	b := 1
// 	piscine.Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "Hello World!"
// 	s = piscine.StrRev(s)
// 	fmt.Println(s)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.BasicAtoi("12345"))
// 	fmt.Println(piscine.BasicAtoi("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi("000000"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.BasicAtoi2("12345"))
// 	fmt.Println(piscine.BasicAtoi2("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi2("012 345"))
// 	fmt.Println(piscine.BasicAtoi2("Hello World!"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Atoi("12345"))
// 	fmt.Println(piscine.Atoi("0000000012345"))
// 	fmt.Println(piscine.Atoi("012 345"))
// 	fmt.Println(piscine.Atoi("Hello World!"))
// 	fmt.Println(piscine.Atoi("+1234"))
// 	fmt.Println(piscine.Atoi("-1234"))
// 	fmt.Println(piscine.Atoi("++1234"))
// 	fmt.Println(piscine.Atoi("--1234"))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := []int{5, 4, 3, 2, 1, 0}
// 	piscine.SortIntegerTable(s)
// 	fmt.Println(s)
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg := 20
// 	fmt.Println(piscine.IterativeFactorial(arg))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg := 20
// 	fmt.Println(piscine.RecursiveFactorial(arg))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IterativePower(4, 3))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.RecursivePower(-4, 3))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	arg1 := 4
// 	fmt.Println(piscine.Fibonacci(arg1))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.Sqrt(4))
// 	fmt.Println(piscine.Sqrt(3))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.IsPrime(5))
// 	fmt.Println(piscine.IsPrime(4))
// }

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.FindNextPrime(5))
// 	fmt.Println(piscine.FindNextPrime(4))
// }

// package main

// import (
// 	"piscine"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	z01.PrintRune(piscine.FirstRune("Hello!"))
// 	z01.PrintRune(piscine.FirstRune("Salut!"))
// 	z01.PrintRune(piscine.FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }


package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}