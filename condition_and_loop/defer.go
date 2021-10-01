package condition_and_loop

import "fmt"

func demoDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func demoStackingDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
