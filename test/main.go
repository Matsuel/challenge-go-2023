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

package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
