package operator

import "fmt"

func logicalOperatorDemo() {
	var c = true
	var d = false
	fmt.Printf("c = %t\n", c)
	fmt.Printf("d = %t\n", d)
	fmt.Printf("c && d is %t\n", c && d)
	fmt.Printf("c || d is %t\n", c || d)
	fmt.Printf("!c is %t\n", !c)
}
