package main

import (
	"fmt"
)

func main() {

	var input float64
	var input2 float64

	fmt.Print("Enter number: ")

	fmt.Scanf("%f %f", &input, &input2)
	fmt.Println(input, input2)
	out := 0.5 * input * input2
	fmt.Println("oun put:", out)
}
