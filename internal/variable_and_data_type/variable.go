package variable_and_data_type

import "fmt"

func VariableDemo() {
	// static variable declaration
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	// dynamic variable declaration
	var z float64 = 20.0
	y := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("z is of type %T\n", z)
	fmt.Printf("y is of type %T\n", y)

	// mixed variable declaration
	var a, b, c = 3, 4, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}
