package operator

import "fmt"

func ArithmeticOperatorDemo() {
	var x = 40
	var y = 20
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)
	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x %% y = %d\n", x%y)
	x++
	fmt.Printf("x++ = %d\n", x)
	y--
	fmt.Printf("y-- = %d\n", y)
}
