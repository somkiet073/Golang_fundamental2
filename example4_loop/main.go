package main

import (
	"fmt"
)

func main() {
	// for
	for i := 0; i < 1; i++ {
		fmt.Println("i=", i)
	}

	// while
	sum := 1
	for sum < 1000 {
		fmt.Println("Sum=", sum)
		sum++
	}

	// foreach
	product := []string{"ProductA", "ProductB", "ProductC"}
	for _, idx := range product {
		fmt.Println(idx)
	}
}
