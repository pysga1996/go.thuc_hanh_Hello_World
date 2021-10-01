package operator

import "fmt"

func miscellaneousOperatorDemo() {
	var f = 21
	fmt.Printf("f = %d\n", f)
	fmt.Printf("Address of f is %p\n", &f)
	var g *int = &f
	fmt.Printf("Pointer p point to address with value %d\n", *g)
}
