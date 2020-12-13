package operator

import "fmt"

func RelationalOperatorDemo() {
	var a = 40
	var b = 20
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("a > b is %t\n", a > b)
	fmt.Printf("a < b is %t\n", b < b)
	fmt.Printf("a == b is %t\n", a == b)
	fmt.Printf("a != b is %t\n", a != b)
	fmt.Printf("a >= b is %t\n", a >= b)
	fmt.Printf("a <= b is %t\n", b <= b)
}
