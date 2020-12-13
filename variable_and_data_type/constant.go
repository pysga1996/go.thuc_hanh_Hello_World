package variable_and_data_type

import "fmt"

func ConstantDemo() {
	// constant
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d", area)
}
