package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Print("Enter number:")
	fmt.Scanf("%d", &input)

	// for
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d \n", input, i, (input * i))
	}

	fmt.Println("< ================= while =================== >")

	// while
	idx := 1
	for idx <= 12 {
		fmt.Printf("%d x %d = %d \n", input, idx, (input * idx))
		idx++
	}
}
