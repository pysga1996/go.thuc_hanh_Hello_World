package condition_and_loop

import (
	"fmt"
	"math"
)

// if demo
func demoIf() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func demoIfWithShortStatement() {
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func demoIfElseWithShortStatement() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
