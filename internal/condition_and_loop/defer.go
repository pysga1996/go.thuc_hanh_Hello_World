package condition_and_loop

import "fmt"

func DemoDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func DemoStackingDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
