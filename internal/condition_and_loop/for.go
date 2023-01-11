package condition_and_loop

import "fmt"

// for demo
func DemoForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// demo for continued loop
func DemoForContinued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// demo for as while loop
func DemoForAsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// demo forever loop

func DemoForever() {
	//for  {
	//
	//}
}
