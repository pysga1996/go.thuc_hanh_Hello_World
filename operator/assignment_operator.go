package operator

import "fmt"

func assignmentOperatorDemo() {
	var e = 60
	fmt.Printf("e = %d\n", e)
	e += 5
	fmt.Printf("e += 5 = %d\n", e)
	e -= 10
	fmt.Printf("e -= 10 = %d\n", e)
	e *= 2
	fmt.Printf("e *= 2 = %d\n", e)
	e /= 11
	fmt.Printf("e /= 11 = %d\n", e)
	e %= 3
	fmt.Printf("e %%= 3 = %d\n", e)
	e <<= 3
	fmt.Printf("e <<= 3 = %d\n", e)
	e >>= 1
	fmt.Printf("e >>= 1 = %d\n", e)
	e &= 2
	fmt.Printf("e ^= 2 = %d\n", e)
	e ^= 2
	fmt.Printf("e |= 2 = %d\n", e)
	e |= 2
	fmt.Printf("e -= 2 = %d", e)
}
