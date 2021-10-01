package operator

import "fmt"

func bitwiseOperatorDemo() {
	var p = 60
	var q = 13
	fmt.Printf("p = %08b (%d)\n", p, p)
	fmt.Printf("q = %08b (%d)\n", q, q)
	fmt.Printf("p & q= %08b (%d)\n", p&q, p&q)
	fmt.Printf("p | q = %08b (%d)\n", p|q, p|q)
	fmt.Printf("p ^ q = %08b (%d)\n", p^q, p^q)
	fmt.Printf("p << 2 %08b (%d)\n", p<<2, p<<2)
	fmt.Printf("p >> 2 %08b (%d)\n", p>>2, p>>2)
}
