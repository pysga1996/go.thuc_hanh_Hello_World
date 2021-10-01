package condition_and_loop

import "fmt"

// for demo
func demoForLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// demo for continued loop
func demoForContinued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// demo for as while loop
func demoForAsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// demo forever loop

func demoForever() {
	//for  {
	//
	//}
}
